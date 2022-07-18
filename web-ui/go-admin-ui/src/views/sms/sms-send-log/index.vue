
<template>
  <BasicLayout>
    <template #wrapper>
      <el-card class="box-card">
        <el-form
          ref="queryForm"
          :model="queryParams"
          :inline="true"
          label-width="68px"
        >
          <el-form-item label="业务编号" prop="phone">
            <el-input
              v-model="queryParams.businessNo"
              placeholder="请输入业务编号"
              clearable
              size="small"
              style="width: 160px"
              @keyup.enter.native="handleQuery"
            />
          </el-form-item>
          <el-form-item label="应用编号" prop="phone">
            <el-input
              v-model="queryParams.appNo"
              placeholder="请输入应用编号"
              clearable
              size="small"
              style="width: 160px"
              @keyup.enter.native="handleQuery"
            />
          </el-form-item>
          <el-form-item>
            <el-button
              type="primary"
              icon="el-icon-search"
              size="mini"
              @click="handleQuery"
              >搜索</el-button
            >
            <el-button icon="el-icon-refresh" size="mini" @click="resetQuery"
              >重置</el-button
            >
          </el-form-item>
        </el-form>

        <el-table
          v-loading="loading"
          :data="smsSendLogList"
          @selection-change="handleSelectionChange"
        >
          <el-table-column type="selection" width="55" align="center" />
          <el-table-column prop="appNo" width="100" label="应用编号" />
          <el-table-column prop="businessNo" width="100" label="业务编号" />
          <el-table-column prop="fee" width="55" label="计价" />
          <el-table-column prop="phoneNumber" width="140" label="发送手机号" />
          <el-table-column
            prop="message"
            label="消息"
            width="120"
            :show-overflow-tooltip="true"
          />
          <el-table-column
            prop="code"
            label="状态码"
            width="100"
            :show-overflow-tooltip="true"
          />
          <el-table-column
            prop="content"
            label="发送内容"
            :show-overflow-tooltip="true"
          />
          <el-table-column
            label="创建时间"
            prop="createdAt"
            sortable="custom"
            width="155"
          >
            <template slot-scope="scope">
              <span>{{ parseTime(scope.row.createdAt) }}</span>
            </template>
          </el-table-column>
          <el-table-column
            prop="remark"
            label="备注"
            width="100"
            :show-overflow-tooltip="true"
          />
        </el-table>

        <pagination
          v-show="total > 0"
          :total="total"
          :page.sync="queryParams.pageIndex"
          :limit.sync="queryParams.pageSize"
          @pagination="getList"
        />

        <!-- 添加或修改对话框 -->
        <el-dialog :title="title" :visible.sync="open" width="500px">
          <el-form ref="form" :model="form" :rules="rules" label-width="80px">
            <el-form-item label="应用ID" prop="appNo">
              <el-input v-model="form.appNo" placeholder="应用ID" />
            </el-form-item>
            <el-form-item label="业务编号" prop="businessNo">
              <el-input v-model="form.businessNo" placeholder="业务编号" />
            </el-form-item>
            <el-form-item label="状态" prop="status">
              <el-input v-model="form.status" placeholder="状态" />
            </el-form-item>
            <el-form-item label="计价条数" prop="fee">
              <el-input v-model="form.fee" placeholder="计价条数" />
            </el-form-item>
            <el-form-item label="发送手机号" prop="phoneNumber">
              <el-input v-model="form.phoneNumber" placeholder="发送手机号" />
            </el-form-item>
            <el-form-item label="接口响应消息" prop="message">
              <el-input v-model="form.message" placeholder="接口响应消息" />
            </el-form-item>
            <el-form-item label="接口响应状态码" prop="code">
              <el-input v-model="form.code" placeholder="接口响应状态码" />
            </el-form-item>
            <el-form-item label="发送内容" prop="content">
              <el-input v-model="form.content" placeholder="发送内容" />
            </el-form-item>
            <el-form-item label="备注" prop="remark">
              <el-input v-model="form.remark" placeholder="备注" />
            </el-form-item>
            <el-form-item label="扩展信息" prop="extJson">
              <el-input v-model="form.extJson" placeholder="扩展信息" />
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
  addSmsSendLog,
  delSmsSendLog,
  getSmsSendLog,
  listSmsSendLog,
  updateSmsSendLog,
} from "@/api/admin/sms-send-log";

export default {
  name: "SmsSendLog",
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
      smsSendLogList: [],

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
      listSmsSendLog(this.addDateRange(this.queryParams, this.dateRange)).then(
        (response) => {
          this.smsSendLogList = response.data.list;
          this.total = response.data.count;
          this.loading = false;
        }
      );
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
        appNo: undefined,
        businessNo: undefined,
        status: undefined,
        fee: undefined,
        phoneNumber: undefined,
        message: undefined,
        code: undefined,
        content: undefined,
        remark: undefined,
        extJson: undefined,
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
      this.title = "添加短信发送记录";
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
      getSmsSendLog(id).then((response) => {
        this.form = response.data;
        this.open = true;
        this.title = "修改短信发送记录";
        this.isEdit = true;
      });
    },
    /** 提交按钮 */
    submitForm: function () {
      this.$refs["form"].validate((valid) => {
        if (valid) {
          if (this.form.id !== undefined) {
            updateSmsSendLog(this.form).then((response) => {
              if (response.code === 200) {
                this.msgSuccess(response.msg);
                this.open = false;
                this.getList();
              } else {
                this.msgError(response.msg);
              }
            });
          } else {
            addSmsSendLog(this.form).then((response) => {
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
          return delSmsSendLog({ ids: Ids });
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
