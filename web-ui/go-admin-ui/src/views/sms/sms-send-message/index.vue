<template>
  <BasicLayout>
    <template #wrapper>
      <el-card class="box-card">
        <div>
          <h1>发送短信测试</h1>
          <el-row :gutter="15">
            <el-form
              ref="formData"
              :model="formData"
              :rules="rules"
              size="medium"
              label-width="100px"
            >
              <el-col :span="12">
                <el-form-item label="业务编码" prop="businessNo">
                  <el-select
                    v-model="formData.businessNo"
                    placeholder="请选择业务编码"
                    clearable
                    :style="{ width: '100%' }"
                    @change="changeGetByTemplateNo"
                  >
                    <el-option
                      v-for="(item, index) in businessNoOptions"
                      :key="index"
                      :label="item.businessName"
                      :value="item.businessNo"
                      :disabled="item.disabled"
                    ></el-option>
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="11">
                <el-form-item label="手机号" prop="phones">
                  <el-input
                    v-model="formData.phones"
                    placeholder="请输入手机号"
                    clearable
                    :style="{ width: '100%' }"
                  >
                  </el-input>
                </el-form-item>
              </el-col>
              <el-col :span="23">
                <el-form-item label="参数" prop="params">
                  <el-input
                    v-model="formData.params"
                    placeholder="请输入参数"
                    clearable
                    :style="{ width: '100%' }"
                  >
                  </el-input>
                </el-form-item>
              </el-col>
              <el-col :span="23">
                <el-form-item label="模版内容" prop="templateContent">
                  <el-input
                    v-model="templateContent"
                    placeholder="请输入参数"
                    readonly
                    clearable
                    :style="{ width: '100%' }"
                  >
                  </el-input>
                </el-form-item>
              </el-col>
              <el-col :span="11">
                <el-form-item label="" prop="field107">
                  <el-button
                    type="primary"
                    icon="el-icon-circle-check"
                    size="medium"
                    @click="submitForm"
                  >
                    发送
                  </el-button>
                </el-form-item>
              </el-col>
              <el-col :span="23">
                <el-form-item label="请求参数" prop="res">
                  <el-input
                    v-model="formData.res"
                    type="textarea"
                    placeholder="请输入响应结果"
                    readonly
                    :autosize="{ minRows: 4, maxRows: 4 }"
                    :style="{ width: '100%' }"
                  ></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="23">
                <el-form-item label="响应结果" prop="resp">
                  <el-input
                    v-model="formData.resp"
                    type="textarea"
                    placeholder="请输入响应结果"
                    readonly
                    :autosize="{ minRows: 4, maxRows: 4 }"
                    :style="{ width: '100%' }"
                  ></el-input>
                </el-form-item>
              </el-col>
            </el-form>
          </el-row>
        </div> </el-card></template
  ></BasicLayout>
</template>
</BasicLayout>
</template>

<script>
import {
  listSmsBusinessConfig,
  getByBusinessNo,
} from "@/api/admin/sms-business-config";
import { getByTemplateNo } from "@/api/admin/sms-template-config";
import { smsSendMessage } from "@/api/admin/sms-send-message";
export default {
  components: {},
  props: [],
  data() {
    return {
      formData: {
        res: "",
        params: "",
        phones: "",
        businessNo: "",
        resp: "",
      },
      templateContent: "",
      rules: {
        businessNo: [
          {
            required: true,
            message: "请选择业务编码",
            trigger: "change",
          },
        ],
        phones: [
          {
            required: true,
            message: "请输入手机号",
            trigger: "blur",
          },
        ],
        params: [
          {
            required: true,
            message: "请输入参数",
            trigger: "blur",
          },
        ],
      },
      businessNoOptions: [],
    };
  },
  created() {
    listSmsBusinessConfig({ pageSize: 1000 }).then((resp) => {
      this.businessNoOptions = resp.data.list;
    });
  },
  methods: {
    submitForm() {
      var params = this.formData.params.split(",");
      var phones = this.formData.phones.split(",");
      var phoneList = [];
      phones.forEach((item) => {
        phoneList.push("+86" + item);
      });
      var data = {
        businessNo: this.formData.businessNo,
        phones: phoneList,
        params: params,
      };
      this.formData.res = JSON.stringify(data);
      smsSendMessage(data).then((resp) => {
        this.formData.resp = JSON.stringify(resp);
      });
    },
    changeGetByTemplateNo() {
      getByBusinessNo(this.formData.businessNo).then((resp) => {
        getByTemplateNo(resp.data.templateNo).then((resp) => {
          this.templateContent = resp.data.templateContent;
        });
      });
    },
    resetForm() {
      this.$refs["elForm"].resetFields();
    },
  },
};
</script>
<style>
</style>
