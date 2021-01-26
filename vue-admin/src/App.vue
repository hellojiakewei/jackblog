<template>
  <div id="nav">
    <router-view v-slot="{ Component,route }">
      <keep-alive>
        <component  :is="Component" />
      </keep-alive>
    </router-view>
  </div>
</template>

<style lang="scss">

</style>