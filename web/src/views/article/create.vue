<template>
  <div class="create-article">
    <el-card shadow="hover">
      <template #header>
        <span>发布动态</span>
      </template>

      <el-form ref="formRef" :model="form" :rules="rules" label-position="top">
        <el-form-item label="标题" prop="title">
          <el-input v-model="form.title" placeholder="说点什么吧..." maxlength="255" show-word-limit />
        </el-form-item>

        <el-form-item label="内容" prop="content">
          <el-input
            v-model="form.content"
            type="textarea"
            :rows="6"
            placeholder="分享你的想法..."
            maxlength="5000"
            show-word-limit
          />
        </el-form-item>

        <el-form-item label="图片/视频">
          <el-upload
            v-model:file-list="fileList"
            :auto-upload="false"
            list-type="picture-card"
            multiple
            :before-upload="() => false"
          >
            <el-icon><Plus /></el-icon>
          </el-upload>
          <div class="upload-tip">支持 jpg、png、gif、mp4，单文件不超过 50MB</div>
        </el-form-item>

        <el-form-item>
          <el-button type="primary" :loading="loading" @click="handleSubmit">
            发布
          </el-button>
          <el-button @click="$router.back()">取消</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import { createArticle } from '@/api/article'
import { ElMessage } from 'element-plus'
import type { FormInstance, FormRules, UploadUserFile } from 'element-plus'

const router = useRouter()
const form = reactive({ title: '', content: '' })
const fileList = ref<UploadUserFile[]>([])
const loading = ref(false)
const formRef = ref<FormInstance>()

const rules: FormRules = {
  title: [{ required: true, message: '请输入标题', trigger: 'blur' }],
  content: [{ required: true, message: '请输入内容', trigger: 'blur' }],
}

async function handleSubmit() {
  const valid = await formRef.value?.validate().catch(() => false)
  if (!valid) return

  loading.value = true
  try {
    const fd = new FormData()
    fd.append('title', form.title)
    fd.append('content', form.content)

    fileList.value.forEach((f) => {
      if (f.raw) {
        fd.append('files', f.raw)
      }
    })

    await createArticle(fd)
    ElMessage.success('发布成功！')
    router.push('/')
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.create-article {
  max-width: 720px;
  margin: 0 auto;
}

.upload-tip {
  font-size: 12px;
  color: #999;
  margin-top: 4px;
}
</style>
