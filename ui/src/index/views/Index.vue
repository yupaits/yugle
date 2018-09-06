<template>
  <div>
    <a-button class="btn-success" @click="test">测试</a-button>
  </div>
</template>

<script>
  export default {
    name: "Index",
    methods: {
      test() {
        this.api.get('/test').then(res => {
          console.log(res);
        });
      }
    }
  }
</script>

<style scoped>

</style>