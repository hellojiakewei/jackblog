<template>
  <div class="header">
    <div class="logo">我是logo</div>
    <div class="personal-center">
      <el-dropdown>
        <i class="el-icon-setting" style="margin-right: 15px"></i>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item>查看</el-dropdown-item>
            <el-dropdown-item>新增</el-dropdown-item>
            <el-dropdown-item>删除</el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
      <span>贾克伟</span>
    </div>
  </div>
</template>

<script>
export default {
  name: "Header",
};
</script>

<style lang="scss" scoped>
.header{
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    .logo{
        width: 180px;
    }
}
</style>