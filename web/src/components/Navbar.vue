<template>
  <div class="logo" style="display: flex; align-items: center;">
    <img style="width: 30px; height: 30px; margin: 10px 5px 5px;" src="../assets/logo.jpg" alt="">
    <el-text :style="{ fontSize: '18px' }">郑航实验室管理系统</el-text>
  </div>

  <el-menu :default-active="activeIndex" class="el-menu-vertical-demo">
    <el-menu-item index="user" @click="goTo('user')">用户管理</el-menu-item>
    <el-menu-item index="students" @click="goTo('students')">学生管理</el-menu-item>
    <el-menu-item index="laboratories" @click="goTo('laboratories')">实验室管理</el-menu-item>
    <el-menu-item index="courses" @click="goTo('courses')">课程管理</el-menu-item>
  </el-menu>
</template>

<script>
export default {
  data() {
    return {
      activeIndex: 'user'
    }
  },
  methods: {
    goTo(page) {
      this.$router.push(`/home/${page}`)
      this.activeIndex = page
    }
  }
}
</script>
