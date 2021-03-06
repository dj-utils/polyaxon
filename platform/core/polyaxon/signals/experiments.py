import logging

from hestia.signal_decorators import ignore_raw, ignore_updates, ignore_updates_pre

from django.db.models.signals import post_save, pre_save
from django.dispatch import receiver

from db.models.experiment_jobs import ExperimentJob
from db.models.experiments import Experiment, ExperimentMetric
from libs.repos.utils import assign_code_reference
from lifecycles.experiments import ExperimentLifeCycle
from lifecycles.jobs import JobLifeCycle
from schemas import ExperimentBackend
from signals.backend import set_backend
from signals.framework import set_framework
from signals.names import set_name
from signals.outputs import set_outputs, set_outputs_refs
from signals.persistence import set_persistence
from signals.tags import set_tags

_logger = logging.getLogger('polyaxon.signals.experiments')


@receiver(pre_save, sender=Experiment, dispatch_uid="experiment_pre_save")
@ignore_updates_pre
@ignore_raw
def experiment_pre_save(sender, **kwargs):
    instance = kwargs['instance']
    # Check if declarations need to be set
    if not instance.params and instance.specification:
        instance.params = instance.specification.params
    set_tags(instance=instance)
    set_persistence(instance=instance)
    set_outputs(instance=instance)
    set_outputs_refs(instance=instance)
    set_name(instance=instance, query=Experiment.all)
    set_backend(instance=instance, default_backend=ExperimentBackend.NATIVE)
    set_framework(instance=instance)
    if not instance.specification or not instance.specification.build:
        return

    if instance.is_independent:
        assign_code_reference(instance)
    else:
        instance.code_reference = instance.experiment_group.code_reference


@receiver(post_save, sender=Experiment, dispatch_uid="experiment_post_save")
@ignore_updates
@ignore_raw
def experiment_post_save(sender, **kwargs):
    instance = kwargs['instance']
    instance.set_status(ExperimentLifeCycle.CREATED)
    if instance.is_independent:
        # TODO: Clean outputs and logs
        pass


@receiver(post_save, sender=ExperimentJob, dispatch_uid="experiment_job_post_save")
@ignore_updates
@ignore_raw
def experiment_job_post_save(sender, **kwargs):
    instance = kwargs['instance']
    instance.set_status(status=JobLifeCycle.CREATED)


@receiver(post_save, sender=ExperimentMetric, dispatch_uid="experiment_metric_post_save")
@ignore_updates
@ignore_raw
def experiment_metric_post_save(sender, **kwargs):
    instance = kwargs['instance']
    experiment = instance.experiment
    experiment.set_metric(metric=instance.values)
