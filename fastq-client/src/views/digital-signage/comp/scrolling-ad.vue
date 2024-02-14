<template>
    <div class="scrolling-ad" :style="{ animationDuration: animationDuration }">
      <span class="text-h5">{{ adText }}</span>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      adText: {
        type: String,
        required: true,
      },
      scrollSpeed: {
        type: Number,
        default: 10, // Adjust as needed
      },
    },
    computed: {
      animationDuration() {
        return `${this.adText.length / this.scrollSpeed}s`;
      },
    },
  };
  </script>
  
  <style scoped>
  .scrolling-ad {
    white-space: nowrap;
    overflow: hidden;
    animation: scroll-left linear infinite;
  }
  
  @keyframes scroll-left {
    to {
      transform: translateX(-100%);
    }
  }
  </style>