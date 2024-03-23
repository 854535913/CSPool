<script setup lang="ts">
import {onMounted, ref} from 'vue'
import {House, Monitor, More} from "@element-plus/icons-vue";
import {getUserInfoService} from "@/api/user";
import useUserInfoStore from '@/store/userInfo.js'
import {logoutService} from "@/api/login";
const userInfoStore = useUserInfoStore();
const handleSelect = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}

onMounted(async () => {
  await getUserInfoService();
});

</script>

<template>
  <div class="common-layout">
    <el-container>
      <el-header>
        <el-menu
            class="el-menu-demo"
            mode="horizontal"
            :ellipsis="false"
            :router="true"
            @select="handleSelect"
        >
          <el-menu-item index="0">
            <h3>CSPool</h3>
          </el-menu-item>
          <div class="flex-grow" />
          <el-menu-item index="/user/login"><span style="font-size: 18px;">Login</span></el-menu-item>
          <el-sub-menu >
            <template #title>
              <span v-if="userInfoStore.info.username" style="font-size: 18px;">{{userInfoStore.info.level}}: {{userInfoStore.info.username}}</span>
              <span v-else style="font-size: 18px;">Visitor</span>
            </template>
            <el-menu-item >Profile</el-menu-item>
            <el-menu-item @click="logoutService()">Logout</el-menu-item>
          </el-sub-menu>
        </el-menu>
      </el-header>
      <el-container>
        <el-aside width="200px">
          <el-menu  class="el-menu-vertical-demo" :router="true" >
            <el-menu-item >
              <el-icon><House /></el-icon>
              <span>Home</span>
            </el-menu-item>
            <el-sub-menu>
              <template #title>
                <el-icon><Monitor /></el-icon>
                <span>CSPool</span>
              </template>
              <el-menu-item index="/post/list">Post list</el-menu-item>
              <el-menu-item index="/post/upload">Create Post</el-menu-item>
            </el-sub-menu>
            <el-sub-menu >
              <template #title>
                <el-icon><More /></el-icon>
                <span>Coming soon</span>
              </template>
              <el-menu-item >module a</el-menu-item>
              <el-menu-item >module b</el-menu-item>
              <el-menu-item >module c</el-menu-item>
            </el-sub-menu>
          </el-menu>
        </el-aside>
        <el-container>
          <el-main>
            <router-view></router-view>
          </el-main>
        </el-container>
      </el-container>
      <el-footer>©2024 Created by paopao</el-footer>
    </el-container>
  </div>
</template>

<style>
.flex-grow {
  flex-grow: 1;
}

.common-layout{
  .el-footer {
    height: 10px;
    text-align: center;
    font-size: 15px;
  }

  .el-container {
    height: 100vh;
  }
}

</style>