<template>
  <div class="home">
    <el-alert
      v-if="!userStore.isLoggedIn"
      title="登录后可以发布动态、评论和添加好友哦"
      type="info"
      show-icon
      :closable="false"
      class="tip-alert"
    >
      <template #default>
        <el-button type="primary" size="small" @click="$router.push('/login')">去登录</el-button>
      </template>
    </el-alert>

    <div class="article-list">
      <el-empty v-if="!loading && articles.length === 0" description="还没有动态，快去发布第一条吧！">
        <el-button v-if="userStore.isLoggedIn" type="primary" @click="$router.push('/create')">
          发布动态
        </el-button>
      </el-empty>

      <div v-for="article in articles" :key="article.id" class="article-card" @click="$router.push(`/article/${article.id}`)">
        <div class="article-header">
          <h3 class="article-title">{{ article.title }}</h3>
        </div>
        <div class="article-content">{{ article.content }}</div>
        <div v-if="article.media?.length" class="article-media">
          <img
            v-for="m in article.media.filter((m: any) => m.type === 'image').slice(0, 4)"
            :key="m.id"
            :src="m.url"
            class="media-thumb"
          />
        </div>
        <div class="article-footer">
          <span class="article-meta">{{ formatDate(article.created_at) }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useUserStore } from '@/stores/user'
import { articleList, type Article } from '@/api/article'

const userStore = useUserStore()
const articles = ref<Article[]>([])
const loading = ref(true)

onMounted(async () => {
  try {
    const res = await articleList()
    articles.value = res.data.articles || []
  } catch {
  } finally {
    loading.value = false
  }
})

function formatDate(dateStr: string) {
  if (!dateStr) return ''
  return new Date(dateStr).toLocaleString('zh-CN')
}
</script>

<style scoped>
.home {
  max-width: 720px;
  margin: 0 auto;
}

.tip-alert {
  margin-bottom: 24px;
}

.article-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.article-card {
  background: #fff;
  border-radius: 8px;
  padding: 20px 24px;
  cursor: pointer;
  transition: box-shadow 0.2s;
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.06);
}

.article-card:hover {
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.1);
}

.article-title {
  font-size: 18px;
  font-weight: 600;
  color: #222;
}

.article-content {
  margin: 12px 0;
  font-size: 15px;
  color: #555;
  line-height: 1.7;
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.article-media {
  display: flex;
  gap: 8px;
  flex-wrap: wrap;
  margin-top: 12px;
}

.media-thumb {
  width: 120px;
  height: 120px;
  object-fit: cover;
  border-radius: 6px;
}

.article-footer {
  margin-top: 16px;
  display: flex;
  align-items: center;
  gap: 16px;
}

.article-meta {
  font-size: 13px;
  color: #999;
}
</style>
