<template>
  <div class="login-container">
    <el-form class="login-form"
             :style="{ height: '200px', padding: '30px'}"
    >
      <el-text
          :style="{ fontSize: '24px', fontWeight: 'bold', margin: '20px', padding: '10px',}"
      >欢迎登录实验室管理信息系统</el-text>
      <el-form-item
        :style="{marginTop: '40px'}"
      >
        <el-input v-model="number" placeholder="工号"></el-input>
      </el-form-item>
      <el-form-item>
        <el-input type="password" v-model="password" placeholder="密码"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="login">登录</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { login } from '@/api/auth';
import { mapActions } from 'vuex';
export default {
  data() {
    return {
      number: '',
      password: ''
    }
  },
  methods: {
    ...mapActions(['setNumber']),
    ...mapActions(['setName']),
    async login() {
      try {
        const response = await login(this.number, this.password);
        if (response.data.message) {
          this.$message.success('登录成功');
          await this.setNumber(this.number);
          await this.setName(response.data.name)
          this.$router.push('/home/user');
        } else {
          this.$message.error('登录失败，请检查用户名和密码。');
        }
      } catch (error) {
        console.error(error);
        this.$message.error('登录出错，请稍后再试。');
      }
    },
  }
}
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background: url('../assets/background.svg') no-repeat center center;
  background-size: cover;
}

.login-form {
  background: rgba(255, 255, 255, 0.8);
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
}
</style>
