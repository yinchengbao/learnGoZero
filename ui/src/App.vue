
<script setup>
import { Button, Notify } from '@opentiny/vue'
import {getPersonInfo } from  '@/axios/api'

const  TinyButton = Button
const click = async () =>  {

  let res = await getPersonInfo({});
  console.log('res');
  console.log(res);

  Notify({type: 'info', title: '自定义标题',message: '获取环境变量：', position: 'top-right'})
}



</script>
<template>
  <tiny-button type="success"  @click="click"> 主要按钮 </tiny-button>
</template>

<style scoped>
</style>
