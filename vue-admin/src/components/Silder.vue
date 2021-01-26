<template>
  <el-menu
    :uniqueOpened="true"
    class="el-menu-vertical-demo"
    background-color="#545c64"
    text-color="#fff"
    active-text-color="#ffd04b"
    router
  >
    <el-submenu v-for="(item, index) in menuList" :index="index.toString()">
      <template #title><i :class="item.icon"></i>{{ item.name }}</template>
      <el-menu-item
        v-for="(options, key) in item.children"
        :index="options.path"
        >{{ options.name }}</el-menu-item
      >
    </el-submenu>
  </el-menu>
</template>

<script lang="ts">
import { defineComponent, reactive } from "vue";
export default defineComponent({
  name: "Silder",
  setup() {
    const menuList = reactive([
      {
        name: "文章",
        icon: "el-icon-message",
        children: [
          { name: "文章发布", path: "/home/articleAdd" },
          { name: "文章列表", path: "/home/articleList" },
        ],
      },
      {
        name: "图片",
        path: "/ddddd",
        icon: "el-icon-picture",
        children: [
          { name: "图片添加", path: "/ddddd" },
          { name: "图片列表", path: "/ddddd" },
        ],
      },
      {
        name: "栏目",
        path: "/category",
        icon: "el-icon-s-operation",
        children: [
          { name: "类型添加", path: "/home/categoryAdd" },
          { name: "类别列表", path: "/home/categoryList" },
        ],
      },
    ]);
    const sequence = (a: number, b: number): string => {
      return `${a}-${b}`;
    };
    return {
      menuList,
      sequence,
    };
  },
});
</script>

<style lang="scss" scoped>
</style>