
<template>
  <BasicLayout>
    <template #wrapper>
      <el-card class="box-card">
        <el-row :gutter="10" class="mb8">
          <el-col :span="1.5">
            <el-button
              v-permisaction="['admin:smsTemplateConfig:add']"
              type="primary"
              icon="el-icon-plus"
              size="mini"
              @click="handleAdd"
              >新增
            </el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
              v-permisaction="['admin:smsTemplateConfig:edit']"
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
              v-permisaction="['admin:smsTemplateConfig:remove']"
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
          :data="smsTemplateConfigList"
          @selection-change="handleSelectionChange"
        >
          <el-table-column type="selection" width="55" align="center" />
          <el-table-column
            prop="templateNo"
            width="100"
            label="模版编号"
            :show-overflow-tooltip="true"
          />
          <el-table-column
            prop="templateName"
            width="100"
            label="模版名称"
            :show-overflow-tooltip="true"
          />
          <el-table-column
            prop="templateContent"
            label="模版内容"
            :show-overflow-tooltip="true"
          />
          <el-table-column
            prop="signName"
            width="100"
            label="签名名称"
            :show-overflow-tooltip="true"
          />
          <el-table-column
            prop="providerNo"
            width="100"
            label="服务商编号"
            :show-overflow-tooltip="true"
          />
          <el-table-column
            prop="thirdPartyTemplateNo"
            width="100"
            label="服务商模版ID"
            :show-overflow-tooltip="true"
          />

          <el-table-column
            prop="remark"
            width="100"
            label="备注"
            :show-overflow-tooltip="true"
          />
          <el-table-column label="状态" width="80" sortable="custom">
            <template slot-scope="scope">
              <el-switch
                v-model="scope.row.status"
                active-value="2"
                inactive-value="1"
                @change="handleStatusChange(scope.row)"
              />
            </template>
          </el-table-column>

          <el-table-column
            label="操作"
            width="150"
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
                  v-permisaction="['admin:smsTemplateConfig:edit']"
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
                  v-permisaction="['admin:smsTemplateConfig:remove']"
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
            <el-form-item label="签名名称" prop="signName">
              <el-select
                v-model="form.signName"
                placeholder="请选择"
                @change="$forceUpdate()"
              >
                <el-option
                  v-for="item in smsSignOptions"
                  :key="item.signName"
                  :label="item.signName"
                  :value="item.signName"
                  :disabled="item.status == 1"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="模版名称" prop="templateName">
              <el-input v-model="form.templateName" placeholder="模版名称" />
            </el-form-item>
            <el-form-item label="模版内容" prop="templateContent">
              <el-input v-model="form.templateContent" placeholder="模版内容" />
            </el-form-item>
            <el-form-item label="备注" prop="remark">
              <el-input v-model="form.remark" placeholder="备注" />
            </el-form-item>
            <el-form-item label="状态" prop="status">
              <el-radio-group v-model="form.status">
                <el-radio
                  v-for="dict in statusOptions"
                  :key="dict.value"
                  :label="dict.value"
                  >{{ dict.label }}</el-radio
                >
              </el-radio-group>
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
  addSmsTemplateConfig,
  delSmsTemplateConfig,
  getSmsTemplateConfig,
  listSmsTemplateConfig,
  updateSmsTemplateConfig,
} from "@/api/admin/sms-template-config";

import { listSmsSignConfig } from "@/api/admin/sms-sign-config";
export default {
  name: "SmsTemplateConfig",
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
      smsTemplateConfigList: [],

      // 状态数据字典
      statusOptions: [],
      //模版签名集合对象
      smsSignOptions: [],

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
    this.getDicts("sys_normal_disable").then((response) => {
      this.statusOptions = response.data;
    });
  },
  methods: {
    /** 查询参数列表 */
    getList() {
      this.loading = true;
      listSmsTemplateConfig(this.queryParams).then((response) => {
        this.smsTemplateConfigList = response.data.list;
        this.total = response.data.count;
        this.loading = false;
      });
    },
    // 取消按钮
    cancel() {
      this.open = false;
      this.reset();
    },
    // 表单重置
    reset() {
      this.form = {
        id: undefined,
        signName: undefined,
        templateNo: undefined,
        templateName: undefined,
        templateContent: undefined,
        thirdPartyTemplateNo: undefined,
        providerNo: undefined,
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
      //获取短信签名集合
      listSmsSignConfig({ pageSize: 1000 }).then((response) => {
        this.smsSignOptions = response.data.list;
      });

      this.title = "添加短信模版配置";
      this.isEdit = false;
    },
    // 模版状态修改
    handleStatusChange(row) {
      const text = row.status === "2" ? "启用" : "停用";
      this.$confirm(
        '确认要"' + text + '""' + row.templateNo + '"模版吗?',
        "警告",
        {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "warning",
        }
      )
        .then(function () {
          return updateSmsTemplateConfig(row);
        })
        .then(() => {
          this.msgSuccess(text + "成功");
        })
        .catch(function () {
          row.status = row.status === "2" ? "1" : "2";
        });
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
      getSmsTemplateConfig(id).then((response) => {
        this.form = response.data;
        this.open = true;
        this.title = "修改短信模版配置";
        this.isEdit = true;
      });
    },
    /** 提交按钮 */
    submitForm: function () {
      this.$refs["form"].validate((valid) => {
        if (valid) {
          if (this.form.id !== undefined) {
            updateSmsTemplateConfig(this.form).then((response) => {
              if (response.code === 200) {
                this.msgSuccess(response.msg);
                this.open = false;
                this.getList();
              } else {
                this.msgError(response.msg);
              }
            });
          } else {
            addSmsTemplateConfig(this.form).then((response) => {
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
          return delSmsTemplateConfig({ ids: Ids });
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
