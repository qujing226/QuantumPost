<template>
  <div style="margin-top: 20px">
    <el-button type="primary" @click="openPasswordDrawer">修改密码</el-button>
    <el-button type="warning" @click="openChangeRequestDrawer">调课申请</el-button>

    <!-- 修改密码的抽屉 -->
    <el-drawer
        title="修改密码"
        v-model="passwordDrawerVisible"
        direction="rtl"
        size="40%"
    >
      <el-input :style="{margin:'20px'}"
          type="password" v-model="newPassword" placeholder="新密码"></el-input>
      <span slot="footer" class="drawer-footer">
        <el-button @click="passwordDrawerVisible = false">取消</el-button>
        <el-button type="primary" @click="savePassword">保存</el-button>
      </span>
    </el-drawer>

    <!-- 调课申请的抽屉 -->
    <el-drawer
        title="调课申请"
        v-model="changeRequestDrawerVisible"
        direction="rtl"
        size="40%"
    >
      <el-form :model="changeRequestForm">
        <el-form-item label="课程">
          <el-input v-model="changeRequestForm.course"></el-input>
        </el-form-item>
        <el-form-item label="申请原因">
          <el-input type="textarea" v-model="changeRequestForm.reason"></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="drawer-footer">
        <el-button @click="changeRequestDrawerVisible = false">取消</el-button>
        <el-button type="primary" @click="submitChangeRequest">提交</el-button>
      </span>
    </el-drawer>
  </div>
</template>

<script>
import { changePassword } from '@/api/auth';
import { requestChangeCourse } from '@/api/course';
import { mapGetters } from 'vuex';
export default {
  computed: { ...mapGetters(['getNumber']) },
  data() {
    return {
      passwordDrawerVisible: false,
      changeRequestDrawerVisible: false,
      newPassword: '',
      changeRequestForm: {
        course: '',
        reason: ''
      },
      number: '',
    }
  },
  created() {
    this.number = this.getNumber;
  },
  methods: {
    openPasswordDrawer() {
      this.passwordDrawerVisible = true;

    },
    async savePassword() {
      console.log(this.newPassword)
      try {
        const response = await changePassword(this.number, this.newPassword);
        if (response.data.message) {
          this.$message.success('密码修改成功');
          this.passwordDrawerVisible = false;
        }else{
          this.$message.error('密码修改失败，请检查新密码。');
        }
      } catch (error) {
        console.error(error);
        this.$message.error('修改密码出错，请稍后再试。');
      }
    },
    openChangeRequestDrawer() {
      this.changeRequestDrawerVisible = true;
    },
    async submitChangeRequest() {
      try {
        await requestChangeCourse(this.changeRequestForm);
        this.$message.success('调课申请已提交');
        this.changeRequestDrawerVisible = false;
      } catch (error) {
        console.error(error);
        this.$message.error('调课申请提交失败，请稍后再试。');
      }
    }
  }
}
</script>

<style scoped>
.drawer-footer {
  text-align: right;
  padding: 20px;
}
</style>
