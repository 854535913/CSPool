<script setup lang="ts">
import {actionPostService, getPostByLikeService, getPostByTimeService} from "@/api/post";
import {onMounted, ref, computed} from "vue";
import {CloseBold, Link, Pointer, Select, Star} from "@element-plus/icons-vue";

// 使用 computed 属性来动态计算最小高度
const minHeight = computed(() => {
  const baseHeight = 100;
  const lineHeight = 18;
  const lines = Math.ceil(post.description.length / 50);
  return baseHeight + lines * lineHeight;
});

const order = ref(1)

const posts_time = ref([]);
const posts_like = ref([]);
onMounted(async () => {
  posts_time.value = await getPostByTimeService();
  posts_like.value = await getPostByLikeService();
});

</script>

<template>
  <div>
    <span style="margin-right: 20px;">Order by: </span>
    <el-radio-group v-model="order">
      <el-radio :value=1 >Time</el-radio>
      <el-radio :value=2 >Like</el-radio>
    </el-radio-group>
  </div>
  <div v-if="order===1">
    <div v-for="post in posts_time" :key="post.title" class="post-container">
      <div class="post-module" :style="{ minHeight: contentHeight + 'px' }">
        <div class="title-and-time">
          <h1 class="post-title">{{ post.title }}</h1>
          <h1 class="post-time">{{ post.postTime }}</h1>
        </div>
        <h1>Author: {{ post.authorName }}</h1>
        <div class="content">{{ post.description }}</div>
        <div>
          <el-link type="info" :href="post.link" :icon="Link" style="font-size: 20px">Link</el-link>
        </div>
        <div class="like">
          <el-icon size="25px"><Star /></el-icon>
          <span>:{{ post.postLike }}</span>
          <el-button :icon="Select" circle style="margin-left: 30px" size="large" @click="actionPostService(post.postID,'like')"/>
          <el-button :icon="CloseBold" circle style="margin-left: 20px" size="large" @click="actionPostService(post.postID,'dislike')"/>
        </div>
      </div>
    </div>
  </div>
  <div v-else>
    <div v-for="post in posts_like" :key="post.title" class="post-container">
      <div class="post-module" :style="{ minHeight: contentHeight + 'px' }">
        <div class="title-and-time">
          <h1 class="post-title">{{ post.title }}</h1>
          <h1 class="post-time">{{ post.postTime }}</h1>
        </div>
        <h1>Author: {{ post.authorName }}</h1>
        <div class="content">{{ post.description }}</div>
        <div>
          <el-link type="info" :href="post.link" :icon="Link" style="font-size: 20px">Link</el-link>
        </div>
        <div class="like">
          <el-icon size="25px"><Star /></el-icon>
          <span>:{{ post.postLike }}</span>
          <el-button :icon="Select" circle style="margin-left: 30px" size="large" @click="actionPostService(post.postID,'like')"/>
          <el-button :icon="CloseBold" circle style="margin-left: 20px" size="large" @click="actionPostService(post.postID,'dislike')"/>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss">

.post-container {
  margin: 30px 60px;
}

.post-module {
  border: 1px solid #ddd;
  border-radius: 8px;
  padding: 20px;
  box-sizing: border-box;
  transition: min-height 0.3s;
}

.title-and-time {
  display: flex;
  justify-content: space-between;
  align-items: center;

  .post-title{
    font-weight: bold;
    font-size: 1.5em;
  }

  .post-time {
    font-size: 0.8em;
    color: #666;
  }
}

.like {
  margin-top: 15px;
  display: flex;
  align-items: center;
  font-size: 1.5em;
}

.post-module h1 {
  margin: 0 0 20px 0;
}

.post-module .content {
  margin-bottom: 20px;
}

.post-module button {
  cursor: pointer;
}

.post-container {
  margin-bottom: 20px;
}

.ml-4{
  color:black;
}
</style>