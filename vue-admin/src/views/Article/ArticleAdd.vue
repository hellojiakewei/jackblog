<template>
  <div>
      <h1>添加文章</h1>
  </div>
</template>

<script>
import {inject} from 'vue'
export default {
    name:"ArticleAdd",
    setup(){
        console.log(inject('name'))
        return {

        }
    }
}
</script>

<style>

</style>