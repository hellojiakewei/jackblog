<template>
  <el-breadcrumb class="app-breadcrumb" separator="/">
    <el-breadcrumb-item v-for="(item, index) in breadcrumbList">
      <a v-if="index != breadcrumbList.length - 1" :href="item.path">{{
        item.name
      }}</a>
      <el-breadcrumb-item v-else>{{ item.name }}</el-breadcrumb-item>
    </el-breadcrumb-item>
  </el-breadcrumb>
</template>

<script>
import { computed, defineComponent, reactive, watch } from "vue";
import { useRouter } from "vue-router";
// interface breadcrumbListProp{
//     name:string,
//     path:string
// }
export default defineComponent({
  name: "BreadCrumb",
  setup() {
    let breadcrumbList = reactive([]);
    const router = useRouter();
    router.currentRoute.value.matched.forEach((item) => {
      breadcrumbList.push({ name: item.meta.name, path: item.path });
    });
    watch(router.currentRoute, () => {
      breadcrumbList.length = 0;
      router.currentRoute.value.matched.forEach((item) => {
        breadcrumbList.push({ name: item.meta.name, path: item.path });
      });
    });
    return {
      breadcrumbList,
    };
  },
});
</script>

<style>
</style>