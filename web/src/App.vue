<template>
  <el-container id="container">
    <el-header id="header">信贷数据检索</el-header>
    <el-container>
      <el-aside id="nav" width="200px">
        <ul>
          <li><router-link to="/loan/query">查询检索</router-link></li>
          <li><router-link to="/loan/export">报表导出</router-link></li>
          <li><router-link to="/loan/update">数据更新</router-link></li>
        </ul>
      </el-aside>
      <el-main id="main">
        <el-scrollbar id="scrollbar" type="height:100%">
          <router-view></router-view>
        </el-scrollbar>
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
export default {
}
</script>

<style>
* {
  margin: 0px;
  padding: 0px;
  font: sans-serif;
}

#container {
  position: fixed;
  width: 100%;
  height: 100%;
}

#header {
  width: 100%;
  height: 150px;
  line-height: 65px;
  background-color: steelblue;
  color: white;
  text-align: center;
  user-select: none;
}

#nav {
  width: 200px;
  height: 100%;
  background-color: lightslategray;
}

#nav a {
  text-decoration: none;
  display: block;
  line-height: 40px;
  padding-left: 40px;
  transition: .2s;
  color: white;
  border-bottom: 1px solid gray;
  user-select: none;
}

#nav a:hover {
  color: red;
  padding-left: 45px;
  border-left: 5px solid red;
}

#main {
  margin: 0px;
  padding: 0px;
  width: 100%;
  height: 100%;
  left: 200px;
  background-color: paleturquoise;
}


</style>
