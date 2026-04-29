<template>
  <div class="article-detail">
    <el-button text @click="$router.push('/')">
      <el-icon><ArrowLeft /></el-icon> 返回
    </el-button>

    <el-card v-if="article" class="detail-card" shadow="hover">
      <h1 class="detail-title">{{ article.title }}</h1>
      <div class="detail-meta">
        <span class="detail-user">用户 #{{ article.user_id }}</span>
        <span class="detail-date">{{ formatDate(article.created_at) }}</span>
        <el-button
          v-if="userStore.userInfo?.id === article.user_id"
          type="danger"
          size="small"
          text
          @click="handleDelete"
        >
          删除
        </el-button>
      </div>

      <div class="detail-content">{{ article.content }}</div>

      <div v-if="article.media?.length" class="detail-media">
        <template v-for="m in article.media" :key="m.id">
          <img v-if="m.type === 'image'" :src="m.url" class="media-img" />
          <video v-else-if="m.type === 'video'" :src="m.url" controls class="media-video" />
        </template>
      </div>
    </el-card>

    <el-empty v-else-if="!loading" description="动态不存在或已删除" />

    <div v-else class="loading-container">
      <el-skeleton :rows="5" animated />
    </div>

    <div v-if="article" class="comment-section">
      <h3 class="comment-title">评论 ({{ comments.length }})</h3>

      <div v-if="userStore.isLoggedIn" class="comment-input">
        <el-input
          v-model="commentText"
          type="textarea"
          :rows="2"
          placeholder="写下你的评论..."
          maxlength="500"
          show-word-limit
        />
        <el-button type="primary" size="small" class="submit-comment" :loading="commenting" @click="handleComment">
          发表
        </el-button>
      </div>

      <div v-if="comments.length" class="comment-list">
        <CommentItem
          v-for="c in comments"
          :key="c.id"
          :comment="c"
          :article-id="articleId"
          @deleted="fetchComments"
          @replied="fetchComments"
        />
      </div>

      <el-empty v-else description="暂无评论" :image-size="80" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'
import { getComments, createComment, deleteArticle, articleDetail } from '@/api/article'
import { ElMessage, ElMessageBox } from 'element-plus'
import CommentItem from './CommentItem.vue'

const route = useRoute()
const router = useRouter()
const userStore = useUserStore()

const articleId = Number(route.params.id)
const article = ref<any>(null)
const comments = ref<any[]>([])
const commentText = ref('')
const commenting = ref(false)
const loading = ref(true)

async function fetchArticle() {
  try {
    const res = await articleDetail(articleId)
    article.value = res.data.article
  } catch {
  } finally {
    loading.value = false
  }
}

async function fetchComments() {
  try {
    const res = await getComments(articleId)
    comments.value = res.data || []
  } catch {}
}

async function handleComment() {
  if (!commentText.value.trim()) return
  commenting.value = true
  try {
    await createComment(articleId, { content: commentText.value })
    commentText.value = ''
    ElMessage.success('评论成功')
    await fetchComments()
  } finally {
    commenting.value = false
  }
}

async function handleDelete() {
  await ElMessageBox.confirm('确定要删除这篇动态吗？', '确认删除', { type: 'warning' })
  try {
    await deleteArticle(articleId)
    ElMessage.success('删除成功')
    router.push('/')
  } catch {}
}

function formatDate(dateStr: string) {
  if (!dateStr) return ''
  return new Date(dateStr).toLocaleString('zh-CN')
}

onMounted(() => {
  fetchArticle()
  fetchComments()
})
</script>

<style scoped>
.article-detail {
  max-width: 720px;
  margin: 0 auto;
}

.detail-card {
  margin: 16px 0 32px;
}

.detail-title {
  font-size: 24px;
  font-weight: 700;
  color: #222;
}

.detail-meta {
  display: flex;
  align-items: center;
  gap: 16px;
  margin: 12px 0 20px;
  font-size: 13px;
  color: #999;
}

.detail-content {
  font-size: 16px;
  line-height: 1.8;
  color: #333;
  white-space: pre-wrap;
}

.detail-media {
  margin-top: 20px;
  display: flex;
  flex-wrap: wrap;
  gap: 12px;
}

.media-img {
  max-width: 100%;
  max-height: 400px;
  border-radius: 8px;
}

.media-video {
  max-width: 100%;
  max-height: 400px;
  border-radius: 8px;
}

.comment-section {
  margin-top: 8px;
}

.comment-title {
  font-size: 18px;
  margin-bottom: 16px;
}

.comment-input {
  margin-bottom: 24px;
}

.submit-comment {
  margin-top: 8px;
}

.comment-list {
  display: flex;
  flex-direction: column;
  gap: 4px;
}
</style>
