<template>
  <div>
    <el-table :data="courses">
      <el-table-column prop="class_time" label="时间"></el-table-column>
      <el-table-column prop="name" label="课程"></el-table-column>
      <el-table-column prop="location" label="地点"></el-table-column>
    </el-table>
  </div>
</template>

<script>
import { getCourses } from '@/api/course'
import {mapGetters} from 'vuex';
export default {
  computed: {...mapGetters(['getNumber', 'getName'])},
  created() {
    this.fetchCourses();
  },
  data() {
    return {
      courses: [
        { day: '星期一', time: '10:00-12:00', course: '计算机科学', location: 'A101' },
        { day: '星期三', time: '14:00-16:00', course: '物理实验', location: 'B202' },
      ]
    }
  },
  methods:{
    async fetchCourses() {
      const number = this.getNumber;
      const response = await getCourses(number);
      if (response.data.data){
        this.courses = response.data.data;
      }else{
        this.$message.error('获取课程失败');
      }
    }
  }
}
</script>

<style scoped>
.el-table {
  width: 100%;
}
</style>
