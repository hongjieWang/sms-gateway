
<template>
  <BasicLayout>
    <template #wrapper>
      <el-card class="box-card">
        <el-row :gutter="10" class="mb8">
          <el-col :span="1.5">
            <el-button
              v-permisaction="['admin:smsServiceProviderConfig:add']"
              type="primary"
              icon="el-icon-plus"
              size="mini"
              @click="handleAdd"
              >新增
            </el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
              v-permisaction="['admin:smsServiceProviderConfig:edit']"
              type="success"
              icon="el-icon-edit"
              size="mini"
              :disabled="single"
              @click="handleUpdate"
              >修改
            </el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
              v-permisaction="['admin:smsServiceProviderConfig:remove']"
              type="danger"
              icon="el-icon-delete"
              size="mini"
              :disabled="multiple"
              @click="handleDelete"
              >删除
            </el-button>
          </el-col>
        </el-row>

        <el-table
          v-loading="loading"
          :data="smsServiceProviderConfigList"
          @selection-change="handleSelectionChange"
        >
          <el-table-column type="selection" width="55" align="center" />
          <el-table-column
            prop="channelNo"
            label="短信渠道商"
            :show-overflow-tooltip="true"
            :formatter="channelNoFormat"
          />
          <el-table-column
            prop="providerNo"
            label="服务商编号"
            :show-overflow-tooltip="true"
          />
          <el-table-column
            prop="providerName"
            label="服务商名称"
            :show-overflow-tooltip="true"
          />
          <el-table-column
            prop="accessKeyId"
            label="身份标识"
            :show-overflow-tooltip="true"
          />
          <el-table-column
            prop="accessKeySecret"
            label="身份认证密钥"
            :show-overflow-tooltip="true"
          />
          <el-table-column
            prop="endpoint"
            label="调用域名"
            :show-overflow-tooltip="true"
          />
          <el-table-column
            prop="sdkAppId"
            label="应用ID"
            :show-overflow-tooltip="true"
          />
          <el-table-column
            prop="region"
            label="地域列表"
            :show-overflow-tooltip="true"
          />

          <el-table-column
            label="操作"
            align="center"
            class-name="small-padding fixed-width"
          >
            <template slot-scope="scope">
              <el-popconfirm
                class="delete-popconfirm"
                title="确认要修改吗?"
                confirm-button-text="修改"
                @onConfirm="handleUpdate(scope.row)"
              >
                <el-button
                  slot="reference"
                  v-permisaction="['admin:smsServiceProviderConfig:edit']"
                  size="mini"
                  type="text"
                  icon="el-icon-edit"
                  >修改
                </el-button>
              </el-popconfirm>
              <el-popconfirm
                class="delete-popconfirm"
                title="确认要删除吗?"
                confirm-button-text="删除"
                @onConfirm="handleDelete(scope.row)"
              >
                <el-button
                  slot="reference"
                  v-permisaction="['admin:smsServiceProviderConfig:remove']"
                  size="mini"
                  type="text"
                  icon="el-icon-delete"
                  >删除
                </el-button>
              </el-popconfirm>
            </template>
          </el-table-column>
        </el-table>

        <pagination
          v-show="total > 0"
          :total="total"
          :page.sync="queryParams.pageIndex"
          :limit.sync="queryParams.pageSize"
          @pagination="getList"
        />

        <!-- 添加或修改对话框 -->
        <el-dialog :title="title" :visible.sync="open" width="800px">
          <el-form ref="form" :model="form" :rules="rules" label-width="100px">
            <el-row>
              <el-col :span="12">
                <el-form-item label="短息渠道商" prop="channelNo">
                  <el-select
                    v-model="form.channelNo"
                    placeholder="请选择"
                    @change="$forceUpdate()"
                  >
                    <el-option
                      v-for="item in channelNoOptions"
                      :key="item.value"
                      :label="item.label"
                      :value="item.value"
                    />
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="服务商编号" prop="providerNo">
                  <el-input
                    v-model="form.providerNo"
                    placeholder="服务商编号"
                  />
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="服务商名称" prop="providerName">
                  <el-input
                    v-model="form.providerName"
                    placeholder="服务商名称"
                  />
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="应用ID" prop="sdkAppId">
                  <el-input
                    v-model="form.sdkAppId"
                    placeholder="应用ID"
                  /> </el-form-item
              ></el-col>
            </el-row>
            <el-form-item label="身份标识" prop="accessKeyId">
              <el-input v-model="form.accessKeyId" placeholder="身份标识" />
            </el-form-item>
            <el-form-item label="身份认证密钥" prop="accessKeySecret">
              <el-input
                v-model="form.accessKeySecret"
                placeholder="身份认证密钥"
              />
            </el-form-item>
            <el-form-item label="调用域名" prop="endpoint">
              <el-input v-model="form.endpoint" placeholder="调用域名" />
            </el-form-item>

            <el-form-item label="地域列表" prop="region">
              <el-input v-model="form.region" placeholder="地域列表" />
            </el-form-item>

            <el-form-item label="状态" prop="status">
              <el-input v-model="form.status" placeholder="状态" />
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button type="primary" @click="submitForm">确 定</el-button>
            <el-button @click="cancel">取 消</el-button>
          </div>
        </el-dialog>
      </el-card>
    </template>
  </BasicLayout>
</template>

<script>
import {
  addSmsServiceProviderConfig,
  delSmsServiceProviderConfig,
  getSmsServiceProviderConfig,
  listSmsServiceProviderConfig,
  updateSmsServiceProviderConfig,
} from "@/api/admin/sms-service-provider-config";

export default {
  name: "SmsServiceProviderConfig",
  components: {},
  data() {
    return {
      // 遮罩层
      loading: true,
      // 选中数组
      ids: [],
      // 非单个禁用
      single: true,
      // 非多个禁用
      multiple: true,
      // 总条数
      total: 0,
      // 弹出层标题
      title: "",
      // 是否显示弹出层
      open: false,
      isEdit: false,
      // 类型数据字典
      typeOptions: [],
      smsServiceProviderConfigList: [],
      channelNoOptions: [
        { label: "腾讯云", value: "Tencent" },
        { label: "阿里云", value: "Aliyun" },
      ],
      // 关系表类型

      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
      },
      // 表单参数
      form: {},
      // 表单校验
      rules: {},
    };
  },
  created() {
    this.getList();
  },
  methods: {
    /** 查询参数列表 */
    getList() {
      this.loading = true;
      listSmsServiceProviderConfig(
        this.addDateRange(this.queryParams, this.dateRange)
      ).then((response) => {
        this.smsServiceProviderConfigList = response.data.list;
        this.total = response.data.count;
        this.loading = false;
      });
    },
    // 取消按钮
    cancel() {
      this.open = false;
      this.reset();
    },
    // 数据状态字典翻译
    channelNoFormat(row, column) {
      return this.selectDictLabel(this.channelNoOptions, row.channelNo);
    },
    // 表单重置
    reset() {
      this.form = {
        id: undefined,
        providerName: undefined,
        providerNo: undefined,
        accessKeyId: undefined,
        accessKeySecret: undefined,
        endpoint: undefined,
        sdkAppId: undefined,
        region: undefined,
        remark: undefined,
        extJson: "{}",
        status: undefined,
      };
      this.resetForm("form");
    },
    getImgList: function () {
      this.form[this.fileIndex] =
        this.$refs["fileChoose"].resultList[0].fullUrl;
    },
    fileClose: function () {
      this.fileOpen = false;
    },
    // 关系
    // 文件
    /** 搜索按钮操作 */
    handleQuery() {
      this.queryParams.pageIndex = 1;
      this.getList();
    },
    /** 重置按钮操作 */
    resetQuery() {
      this.dateRange = [];
      this.resetForm("queryForm");
      this.handleQuery();
    },
    /** 新增按钮操作 */
    handleAdd() {
      this.reset();
      this.open = true;
      this.title = "添加服务商配置";
      this.isEdit = false;
    },
    // 多选框选中数据
    handleSelectionChange(selection) {
      this.ids = selection.map((item) => item.id);
      this.single = selection.length !== 1;
      this.multiple = !selection.length;
    },
    /** 修改按钮操作 */
    handleUpdate(row) {
      this.reset();
      const id = row.id || this.ids;
      getSmsServiceProviderConfig(id).then((response) => {
        this.form = response.data;
        this.open = true;
        this.title = "修改服务商配置";
        this.isEdit = true;
      });
    },
    /** 提交按钮 */
    submitForm: function () {
      this.$refs["form"].validate((valid) => {
        if (valid) {
          if (this.form.id !== undefined) {
            updateSmsServiceProviderConfig(this.form).then((response) => {
              if (response.code === 200) {
                this.msgSuccess(response.msg);
                this.open = false;
                this.getList();
              } else {
                this.msgError(response.msg);
              }
            });
          } else {
            addSmsServiceProviderConfig(this.form).then((response) => {
              if (response.code === 200) {
                this.msgSuccess(response.msg);
                this.open = false;
                this.getList();
              } else {
                this.msgError(response.msg);
              }
            });
          }
        }
      });
    },
    /** 删除按钮操作 */
    handleDelete(row) {
      var Ids = (row.id && [row.id]) || this.ids;

      this.$confirm('是否确认删除编号为"' + Ids + '"的数据项?', "警告", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(function () {
          return delSmsServiceProviderConfig({ ids: Ids });
        })
        .then((response) => {
          if (response.code === 200) {
            this.msgSuccess(response.msg);
            this.open = false;
            this.getList();
          } else {
            this.msgError(response.msg);
          }
        })
        .catch(function () {});
    },
  },
};
</script>
