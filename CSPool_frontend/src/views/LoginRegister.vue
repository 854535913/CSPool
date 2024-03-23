<script setup lang="ts">
import {clearLoginData, loginData, loginRules, loginService} from "@/api/login";
import {ElButton, ElForm, ElInput} from "element-plus";
import {clearRegisterData, registerData, registerRules, registerService} from "@/api/register";
import { ref } from 'vue'
import {Promotion, UploadFilled} from "@element-plus/icons-vue";

const switchLoginRegister = ref(true)
</script>

<template>
  <!-- 登录表单 -->
  <div v-if="switchLoginRegister">
    <h1 class="word1">Enter your username and password to sign in to your CSPool account.</h1>
    <h1 class="word2">Not signed up yet?
      <a href="#" @click="switchLoginRegister = false; clearLoginData(); $event.preventDefault()">Sign up here</a>
    </h1>
    <div class="box" id="login" >
      <el-form :model="loginData" :rules="loginRules">
        <h1 class="word3">Username</h1>
        <el-form-item prop="username">
          <el-input v-model="loginData.username" style="width: 300px;height: 50px" placeholder="Please input username"></el-input>
        </el-form-item>
        <h1 class="word3">Password</h1>
        <el-form-item prop="password">
          <el-input v-model="loginData.password" style="width: 300px;height: 50px" type="password" placeholder="Please input password" show-password/>
        </el-form-item>
        <el-form-item>
          <el-button :icon="Promotion" style="width: 120px;height: 40px;background-color: #333333;color: white" @click="loginService">Login</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>

  <!-- 注册表单 -->
  <div v-else >
    <h1 class="word1">Free registration on CSPool.</h1>
    <h1 class="word2">Already have an account?
      <a href="#" @click="switchLoginRegister = true; clearRegisterData(); $event.preventDefault()">Sign in</a>
    </h1>
    <div class="box" id="register">
      <el-form :model="registerData" :rules="registerRules">
        <h1 class="word3">Username</h1>
        <el-form-item prop="username">
          <el-input v-model="registerData.username"  style="width: 300px;height: 50px" placeholder="Please input username"></el-input>
        </el-form-item>
        <h1 class="word3">Password</h1>
        <el-form-item prop="password">
          <el-input v-model="registerData.password" style="width: 300px;height: 50px" type="password" placeholder="Please input password" show-password/>
        </el-form-item>
        <h1 class="word3">Re_Password</h1>
        <el-form-ite prop="re_password">
          <el-input v-model="registerData.re_password" type="password" style="width: 300px;height: 50px" placeholder="Please input password again" show-password/>
        </el-form-ite>
        <h1 class="word3">Invitation code</h1>
        <el-form-item prop="invitation">
          <el-input v-model="registerData.invitation"  style="width: 300px;height: 50px" placeholder="Invitation code (not required)"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button :icon="UploadFilled" style="width: 120px;height: 40px;background-color: #333333;color: white" @click="registerService">Register</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>


  <!-- 公共组件 -->
  <div>
    <h1 class="welcome">Welcome back</h1>
    <div class="line"></div>
  </div>

</template>

<style lang="scss">

.box {position: absolute;
  width: 530px;
  border: 3px solid var(--el-border-color);
  border-radius: 10px;
  display: flex; /* 启用Flex布局 */
  flex-direction: column; /* 子元素垂直排列 */
  justify-content: center; /* 垂直居中 */
  align-items: center; /* 水平居中 */
}
#login {
  top: 180px;
  left: 650px;
  height: 350px;
}
#register {
  top: 120px;
  left: 650px;
  height: 550px;
}
.welcome {position: absolute;
  font-family: Arial, sans-serif;
  top: 200px;
  left: 100px;
  font-size: 60px;
}

.word1{position: absolute;
  max-width: 500px;
  font-family: Arial, sans-serif;
  top: 320px;
  left: 100px;
  font-size: 20px;
}

.word2{position: absolute;
  max-width: 500px;
  font-family: Arial, sans-serif;
  top: 400px;
  left: 100px;
  font-size: 20px;
}

.word3{
  font-family: Arial, sans-serif;
  font-size: 20px;
}

.link{position: absolute;
  max-width: 500px;
  font-family: Arial, sans-serif;
  top: 405px;
  left: 300px;
  font-size: 20px;
}

.line {position: absolute;
  top: 375px;
  left: 100px;
  border-top: 2px solid #000;
  width: 35%;
  margin: 20px 0;
}

</style>