<script setup lang="ts">
import {ElInput, ElMessage, ElButton, ElForm} from 'element-plus';
import {registerData, registerRules, registerService} from '@/api/register'
</script>

<template>
  <el-row class="login-page">
    <el-col :span="12" class="bg"></el-col>
    <el-col :span="6" :offset="3" class="form">
      <el-form :model="registerData" label-width="auto" style="max-width: 600px" :rules="registerRules">
        <el-form-item>
          <h1>Register</h1>
        </el-form-item>
        <el-form-item prop="username">
          <el-input v-model="registerData.username" placeholder="Please input username"></el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input
              v-model="registerData.password"
              type="password"
              placeholder="Please input password"
              show-password
          />
        </el-form-item>
        <el-form-item prop="re_password">
          <el-input
              v-model="registerData.re_password"
              type="password"
              placeholder="Please input password again"
              show-password
          />
        </el-form-item>
        <el-form-item>
          <el-button class="button" type="primary" @click="registerService">Register</el-button>
        </el-form-item>
      </el-form>
    </el-col>
  </el-row>
</template>

<style scoped lang="scss">
.login-page {
  height: 100vh;
  background-color: #fff;

  .bg {
    background: url('../assets/logo2.png') no-repeat 50% 45% / 240px auto,
    url('../assets/login_bg.jpg') no-repeat center / cover;
    border-radius: 0 20px 20px 0;
  }

  .form {
    display: flex;
    flex-direction: column;
    justify-content: center;
    user-select: none;

    .title {
      margin: 0 auto;
    }

    .button {
      width: 100%;
    }
  }
}
</style>