<template>
  <div class="cp" ref="colorPickerContainer">
      <img class="image" @click="$emit('endColor')" :src="imgUrl">
    <canvas class="canvas" ref="canvas" @mousemove="pickColor"  />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
defineProps<{
  imgUrl?: string
}>()

const canvas = ref<HTMLCanvasElement | null>(null);
const colorPickerContainer = ref<HTMLElement | null>(null);
const onImageLoad = () => {
  if (canvas.value && colorPickerContainer.value) {
    const img = colorPickerContainer.value.querySelector('img') as HTMLImageElement;
    canvas.value.width = img.clientWidth;
    canvas.value.height = img.clientHeight;
    const ctx = canvas.value.getContext('2d');
    if (ctx) {
      ctx.drawImage(img, 0, 0, canvas.value.width, canvas.value.height);
    }
  }
};

const pickColor = (event: MouseEvent) => {
  if (canvas.value) {
    const rect = canvas.value.getBoundingClientRect();
    const x = event.clientX - rect.left;
    const y = event.clientY - rect.top;
    const ctx = canvas.value.getContext('2d');
    if (ctx) {
      const pixel = ctx.getImageData(x, y, 1, 1).data;
      const rgba = `rgba(${pixel[0]}, ${pixel[1]}, ${pixel[2]}, ${pixel[3] / 255})`;
      console.log(rgba)
    }
  }
};

onMounted(() => {
  if (colorPickerContainer.value) {
    const img = colorPickerContainer.value.querySelector('img');
    if (img) {
      img.addEventListener('load', onImageLoad);
    }
  }
});
</script>

<style scoped lang="scss">
.cp{
  .image{
    display: block;
    width: 100vw;
    height: 100vh;
  }
  .canvas {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    pointer-events: auto;
    opacity: 0; /* 让 canvas 透明，但仍然能够接收事件 */
  }
}

</style>
