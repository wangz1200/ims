<template>
  <div id="input">
      <el-input 
        id="text"
        type="textarea"
        placeholder="请输入公式..." 
        v-model="content" 
        :autosize="{ minRows: 1, maxRows: 15}" 
        clearable>
      </el-input><el-button id="button" @click="query">查询</el-button>
  </div>
</template>

<script>
export default {
  name: 'LoanQuery',
  data() {
    return {
      content: "",
    }
  },
  methods: {
    query() {

    }
  }
}
</script>

<style scoped>
#input {
  min-width: 300px;
  padding: 10px;
  box-sizing: border-box;
  display: inline-block;
}
#text {
  display: inline-block;
}
#button {
  display: inline;
}
</style>

