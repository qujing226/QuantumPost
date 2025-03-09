<template>
  <div>
    <el-button type="primary">添加学生</el-button>
    <el-table :data="students">
      <el-table-column prop="id" label="学号"></el-table-column>
      <el-table-column prop="name" label="姓名"></el-table-column>
      <el-table-column label="带队老师">    <template #default>
        <span>{{ name }}</span>
      </template></el-table-column>
      <el-table-column label="操作">
        <template #default="scope">
          <el-button @click="studentEditVisible = true">编辑</el-button>
          <el-dialog
              title="编辑学生信息"
              v-model="studentEditVisible"
              direction="rtl"
              size="40%"
          >
            <el-input :style="{margin:'20px'}"
                      type="password" v-model="newStudentNumber" placeholder="学号"></el-input>
            <span slot="footer" class="drawer-footer">
        <el-button @click="studentEditVisible = false">取消</el-button>
        <el-button type="primary" @click="editStudent(scope.row)">保存</el-button>
      </span>
          </el-dialog>
          <el-button @click="deleteStudent(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import {mapGetters} from 'vuex';
import {getStudents} from "@/api/students";

export default {
  computed: {...mapGetters(['getNumber', 'getName'])},
  data() {
    return {
      students: [{id: '1001', name: '张三', lab: '06-A5'},{id: '1002', name: '李四', lab: '06-A3'}],
      number: '',
      name:'',
      studentEditVisible: false,
      newStudentNumber:"",
    }
  },
  created() {
    this.number = this.getNumber;
    this.name = this.getName;
    this.fetchStudents();
  },
  methods: {
    async fetchStudents() {
      try {
        const response = await getStudents(this.number);
        if (response.data.data) {
          this.students = response.data.data;
        } else {
          this.$message.error('获取学生列表失败');
        }
      } catch (error) {
        console.error(error);
        this.$message.error('获取学生列表失败');
      }
    },
    editStudent(student) {
      // 编辑学生逻辑
    },
    deleteStudent(student) {
      // 删除学生逻辑
    }
  }
}
</script>
