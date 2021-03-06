// Copyright 2019 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * OpenAPI spec version: 1.0.0
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 */


package io.swagger.client.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

/**
 * V1Mapping
 */

public class V1Mapping {
  @SerializedName("kind")
  private String kind = null;

  @SerializedName("values")
  private String values = null;

  @SerializedName("concurrency")
  private Integer concurrency = null;

  @SerializedName("early_stopping")
  private List<Object> earlyStopping = null;

  public V1Mapping kind(String kind) {
    this.kind = kind;
    return this;
  }

   /**
   * Get kind
   * @return kind
  **/
  @ApiModelProperty(value = "")
  public String getKind() {
    return kind;
  }

  public void setKind(String kind) {
    this.kind = kind;
  }

  public V1Mapping values(String values) {
    this.values = values;
    return this;
  }

   /**
   * Get values
   * @return values
  **/
  @ApiModelProperty(value = "")
  public String getValues() {
    return values;
  }

  public void setValues(String values) {
    this.values = values;
  }

  public V1Mapping concurrency(Integer concurrency) {
    this.concurrency = concurrency;
    return this;
  }

   /**
   * Get concurrency
   * @return concurrency
  **/
  @ApiModelProperty(value = "")
  public Integer getConcurrency() {
    return concurrency;
  }

  public void setConcurrency(Integer concurrency) {
    this.concurrency = concurrency;
  }

  public V1Mapping earlyStopping(List<Object> earlyStopping) {
    this.earlyStopping = earlyStopping;
    return this;
  }

  public V1Mapping addEarlyStoppingItem(Object earlyStoppingItem) {
    if (this.earlyStopping == null) {
      this.earlyStopping = new ArrayList<Object>();
    }
    this.earlyStopping.add(earlyStoppingItem);
    return this;
  }

   /**
   * Get earlyStopping
   * @return earlyStopping
  **/
  @ApiModelProperty(value = "")
  public List<Object> getEarlyStopping() {
    return earlyStopping;
  }

  public void setEarlyStopping(List<Object> earlyStopping) {
    this.earlyStopping = earlyStopping;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1Mapping v1Mapping = (V1Mapping) o;
    return Objects.equals(this.kind, v1Mapping.kind) &&
        Objects.equals(this.values, v1Mapping.values) &&
        Objects.equals(this.concurrency, v1Mapping.concurrency) &&
        Objects.equals(this.earlyStopping, v1Mapping.earlyStopping);
  }

  @Override
  public int hashCode() {
    return Objects.hash(kind, values, concurrency, earlyStopping);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1Mapping {\n");
    
    sb.append("    kind: ").append(toIndentedString(kind)).append("\n");
    sb.append("    values: ").append(toIndentedString(values)).append("\n");
    sb.append("    concurrency: ").append(toIndentedString(concurrency)).append("\n");
    sb.append("    earlyStopping: ").append(toIndentedString(earlyStopping)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(java.lang.Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}

