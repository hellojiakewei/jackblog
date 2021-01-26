<template>
  <BreadCrumb />
  <router-view :key="key"> </router-view>
</template>

<script lang="ts">
import BreadCrumb from "@/components/BreadCrumb.vue";
import { computed, defineComponent, watch } from "vue";
import { useRouter } from "vue-router";
export default defineComponent({
  name: "Main",
  components: {
    BreadCrumb,
  },
  setup() {
    const router = useRouter();
    const key = computed(() => {
      return router.currentRoute.value.path;
    });
    return {
      key,
    };
  },
});
</script>

<style lang="scss" scoped>
</style>