<template>
  <div class="common-layout">
    <el-container>
      <el-header>
        <Header />
      </el-header>
      <el-container>
        <el-aside width="200px">
          <Silder />
        </el-aside>
        <el-main>
          <Main />
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script>
import Header from "../components/Header.vue";
import Silder from "../components/Silder.vue";
import Main from "../components/Main.vue";
import {provide} from 'vue'
export default {
  name: "Layout",
  components: {
    Header,
    Silder,
    Main,
  },
  setup(){
      provide('name','jiakewei')
  }
};
</script>
<style lang="scss">
.el-header {
  background-color: rgb(67, 74, 80);
  color: #333;
  line-height: 60px;
}

.el-aside {
  height: 100vh;
  background-color: #545c64;
  text-align: left;
}
</style>