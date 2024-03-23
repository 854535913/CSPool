<script setup lang="ts">
import {ElButton, ElForm} from 'element-plus';
import {postData,uploadService} from "@/api/post";
import {ref} from "vue";
import {Promotion} from "@element-plus/icons-vue";
import {loginService} from "@/api/login";
const checked1 = ref(false)
const checked2 = ref(false)
</script>

<template>
  <h1 class="Createapost">Create a post</h1>
  <div class="post-module" >
    <el-form :model="postData" label-width="auto" style="max-width: 600px">
      <h1 class="post_subtitle">Title</h1>
      <el-form-item >
        <el-input v-model="postData.title" />
      </el-form-item>
      <h1 class="post_subtitle">Description</h1>
      <el-form-item >
        <el-input v-model="postData.description" type="textarea" />
      </el-form-item>
      <div>
        <h1><span style="margin-right: 20px;">Options:</span><el-checkbox v-model="checked1" label="Link" size="large" />
          <el-checkbox v-model="checked2" label="Tag" size="large" /></h1>
      </div>
      <div v-if="checked1">
        <h1 class="post_subtitle">Link</h1>
        <el-form-item >
          <el-input v-model="postData.link" />
        </el-form-item>
      </div>
      <div v-if="checked2">
        <h1 class="post_subtitle">Tag</h1>
        <el-form-item >
          <el-input v-model="postData.tag" />
        </el-form-item>
      </div>
      <div id="post_button">
        <el-form-item>
          <el-button :icon="Promotion" style="width: 120px;height: 40px;background-color: #333333;color: white" @click="uploadService">Post</el-button>
        </el-form-item>
      </div>
    </el-form>
  </div>
</template>


<style scoped lang="scss">

.post-module {
  margin: 30px 200px;
  border: 1px solid #ddd;
  border-radius: 8px;
  padding: 50px;
  box-sizing: border-box;
  transition: min-height 0.3s;
}

.Createapost{
  font-size: 25px;
}

.post_subtitle{
  font-size: 17px;
}

#post_button{
  margin-top: 30px;
}
</style>