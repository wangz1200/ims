<template>
<div>
  <div id="input">
    <div id="text">
      <el-input type="textarea" placeholder="请输入公式..." v-model="content" :autosize="{ minRows: 1, maxRows: 15}" />
    </div>
    <div id="button">
      <el-button id="button" @click="query">查询</el-button>
    </div>
  </div>
  <div id="show" v-show="res">
    show content
  </div>
</div>
</template>

<script>
import axios from "../api"

export default {
  name: 'LoanQuery',
  data() {
    return {
      content: "",
      res: "",
    }
  },
  methods: {
    query() {
      axios.post("/query").then(res=>{
        console.log(res);
      })
    }
  }
}
</script>

<style scoped>
#input {
  padding: 10px;
  box-sizing: border-box;
  display: table-row;
}

#text {
  width: 100%;
  display: table-cell;
  vertical-align: middle;
}

#button {
  display: table-cell;
  vertical-align: middle;
}

</style>

