<template>
  <div class="container">
    <div class="sidebar">
      <navbar/>
    </div>
    <div class="content">
      <div class="welcome">
        <hr>
        <el-text :style="{fontSize: '24px', fontWeight: 'bold', margin: '20px', padding: '10px'}">欢迎您：{{name}}
        </el-text>
        <br>
        <el-text :style="{margin: '20px', padding: '10px',}">{{ welcomeMessage }}</el-text>
        <hr>
      </div>

      <el-carousel v-if="showCarousel" :interval="4000"  style="width: 700px;justify-content: center;  align-items: center;">
        <el-carousel-item v-for="image in images" :key="image">
          <img :src="image" alt="实验室图片" class="carousel-image"/>
        </el-carousel-item>
      </el-carousel>
      <router-view @route-change="updateShowCarousel"/>
    </div>
  </div>
</template>

<script>
import Navbar from '../components/Navbar.vue'
import {mapGetters} from "vuex";

export default {
  computed: {...mapGetters(['getName'])},
  components: {
    Navbar
  },
  data() {
    return {
      images: [
        'https://th.bing.com/th/id/OIP.Fvcr09_SI6adnARhSZf6xwHaEI?rs=1&pid=ImgDetMain',
        'https://img95.699pic.com/photo/40186/7991.jpg_wh860.jpg',
        'https://th.bing.com/th/id/OIP.KE7FtiSEKPcE1fUlxST8bQHaEb?w=506&h=303&rs=1&pid=ImgDetMain'
      ],
      welcomeMessage: '今日有课：计算机系统，A101，10:00-12:00',
      name: '',
      showCarousel: true
    }
  },
  methods: {
    updateShowCarousel() {
      const route = this.$route.path;
      this.showCarousel = route === '/home' || route === '/home/user';
    }
  }, watch: {
    $route(to, from) {
      this.updateShowCarousel();
    }
  },
  created() {
    this.name = this.getName;
    this.updateShowCarousel();
  }
}
</script>

<style scoped>
.container {
  display: flex;
  height: 40vh;
}
.carousel-container {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 40vh; /* 可根据需要调整 */
}
.carousel-image {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  object-fit: cover; /* 保持图片比例并裁剪 */
}
.sidebar {
  width: 250px;
  background-color: #f5f5f5;
  border-right: 1px solid #ddd;
}

.content {
  flex-grow: 1;
  padding: 60px;
}

.welcome {
  height: 100px;
  margin-bottom: 20px;
  font-size: 18px;
  font-weight: bold;
}

.carousel-image {
  width: 100%;
  object-fit: cover;
}
</style>
