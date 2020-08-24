<template>
  <div id="app">
    <div class="siderbar">
      <header class="header">贷款数据</header>
      <ul>
        <li><router-link to="/loan/query">数据查询</router-link></li>
        <li><router-link to="/loan/export">报表导出</router-link></li>
        <li><router-link to="/loan/update">数据更新</router-link></li>
      </ul>
    </div>
    <div class="main">
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
export default {
  name: 'App',
  components: {
  }
}
</script>

<style>

#app {
  position: fixed;
  display: inline-block;
  padding: 0px;
  margin: 0px;
  list-style: none;
  text-decoration: none;
  font-family: 'Roboto', sans-serif;
}

.container {
    height: 100%;
    display: inline-block;
}

.siderbar {
    position: fixed;
    display: block;
    width: 248px;
    height: 100%;
    background-color: #042331;
    border-right: 2px solid black;
}

.header {
    width: 100%;
    text-align: center;
    font-size: 25px;
    line-height: 65px;
    color: white;
    user-select: none;
    background-color: #063146;
}

.siderbar li {
    display: block;
    line-height: 40px;
    user-select: none;
    border-top: 1px solid black;
    border-bottom: 1px solid black;  
}
.siderbar a {
    display: block;
    width: 100%;
    text-decoration: none;
    padding-left: 40px;
    color: white;
    display: block;
    transition: .3s; 
}

.siderbar a:hover {
    padding-left: 50px;
    border-left: 5px solid #b93632; 
    color: #b93632;
}

.router-link-active {
    text-decoration: none;
}

.main {
    position: fixed;
    padding-left: 3px;
    left: 250px;
    width: 100%;
    height: 100%;
}
</style>
