<template>
  <div class="comment-item" :class="{ 'is-reply': comment.parent_id }">
    <div class="comment-body">
      <span class="comment-user">用户#{{ comment.user_id }}</span>
      <span class="comment-time">{{ formatDate(comment.created_at) }}</span>

      <p class="comment-content">{{ comment.content }}</p>

      <div class="comment-actions">
        <el-button
          v-if="userStore.isLoggedIn"
          size="small"
          text
          @click="showReply = !showReply"
        >
          <el-icon><ChatLineSquare /></el-icon>回复
        </el-button>
        <el-button
          v-if="userStore.userInfo?.id === comment.user_id"
          size="small"
          text
          type="danger"
          @click="handleDelete"
        >
          删除
        </el-button>
      </div>

      <div v-if="showReply" class="reply-input">
        <el-input
          v-model="replyText"
          size="small"
          type="textarea"
          :rows="2"
          placeholder="回复..."
          maxlength="500"
        />
        <el-button size="small" type="primary" :loading="replying" @click="handleReply">
          回复
        </el-button>
      </div>
    </div>

    <div v-if="comment.replies?.length" class="replies">
      <CommentItem
        v-for="child in comment.replies"
        :key="child.id"
        :comment="child"
        :article-id="articleId"
        @deleted="$emit('deleted')"
        @replied="$emit('replied')"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useUserStore } from '@/stores/user'
import { createComment, deleteComment } from '@/api/article'
import { ElMessage, ElMessageBox } from 'element-plus'

const props = defineProps<{
  comment: any
  articleId: number
}>()

const emit = defineEmits<{
  deleted: []
  replied: []
}>()

const userStore = useUserStore()
const showReply = ref(false)
const replyText = ref('')
const replying = ref(false)

async function handleReply() {
  if (!replyText.value.trim()) return
  replying.value = true
  try {
    await createComment(props.articleId, {
      content: replyText.value,
      parent_id: props.comment.id,
    })
    replyText.value = ''
    showReply.value = false
    ElMessage.success('回复成功')
    emit('replied')
  } finally {
    replying.value = false
  }
}

async function handleDelete() {
  await ElMessageBox.confirm('确定要删除这条评论吗？', '确认删除', { type: 'warning' })
  try {
    await deleteComment(props.comment.id)
    ElMessage.success('删除成功')
    emit('deleted')
  } catch {}
}

function formatDate(dateStr: string) {
  if (!dateStr) return ''
  return new Date(dateStr).toLocaleString('zh-CN')
}
</script>

<style scoped>
.comment-item {
  padding: 12px 0;
  border-bottom: 1px solid #f0f0f0;
}

.is-reply {
  margin-left: 32px;
  border-left: 2px solid #eee;
  padding-left: 16px;
}

.comment-user {
  font-size: 14px;
  font-weight: 600;
  color: #333;
}

.comment-time {
  font-size: 12px;
  color: #bbb;
  margin-left: 12px;
}

.comment-content {
  margin: 6px 0 8px;
  font-size: 14px;
  line-height: 1.6;
  color: #444;
}

.comment-actions {
  display: flex;
  gap: 4px;
}

.reply-input {
  margin-top: 8px;
  display: flex;
  gap: 8px;
  align-items: flex-end;
}

.replies {
  margin-top: 4px;
}
</style>
