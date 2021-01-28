<template></template>

<script>
import { defineComponent } from "vue";
import { useStore } from "vuex";
import axios from "axios";
export default defineComponent({
  name: "Login",
  setup() {
    //初始化数据
    // const store = useStore();
    // const res = store.dispatch("login", {
    //   username: "jiakewei",
    //   password: "123456",
    // });
    // console.log(res);
    var myHeaders = new Headers();
myHeaders.append("Accept", "application/json, text/plain, */*");
myHeaders.append("Referer", "http://localhost:7777/");
myHeaders.append("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36");
myHeaders.append("Content-Type", "application/json;charset=UTF-8");

var raw = "{\"username\":\"jiakewei\",\"password\":\"123456\"}";

var requestOptions = {
  method: 'POST',
  headers: myHeaders,
  body: raw,
  redirect: 'follow'
};

fetch("http://localhost:8081/v1/user/login", requestOptions)
  .then(response => response.text())
  .then(result => console.log(result))
  .catch(error => console.log('error', error));

    // axios
    //   .post("http://localhost:8081/v1/user/login", {
    //     username:'jiakewei' ,
    //     password: "123456",
    //   })
    //   .then(function (response) {
    //     console.log(response);
    //   })
    //   .catch(function (error) {
    //     console.log(error);
    //   });
  },
});
</script>

<style>
</style>