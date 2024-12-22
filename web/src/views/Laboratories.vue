<template>
  <div>
    <el-table :data="labs">
      <el-table-column prop="id" label="实验室编号"></el-table-column>
      <el-table-column prop="capacity" label="容量"></el-table-column>
      <el-table-column prop="location" label="地点"></el-table-column>
      <el-table-column label="老师">{{name}}</el-table-column>
    </el-table>
  </div>
</template>

<script>
import { getLabs } from '@/api/labs'
import {mapGetters} from 'vuex';
export default {
  computed: {...mapGetters(['getNumber', 'getName'])},
  created() {
    this.number = this.getNumber;
    this.name = this.getName;
    this.fetchLabs();
  },
  data() {
    return {
      labs: [
        { id: '06-A5', studentCount: 30, status: '有课' },
        { id: '06-A3', studentCount: 15, status: '空闲' }
      ],
      number:'',
      name:''
    }
  },
  methods: {
    async fetchLabs() {
      const response =  await getLabs(this.number);
      if (response.data.data){
        this.labs = response.data.data;
      }else{
        this.$message.error('获取实验室列表失败');
      }
    }
  },
}
</script>
