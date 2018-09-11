<template>
  <div>
    <h2>必应壁纸</h2>
    <a-row :gutter="16">
      <a-col :span="8" v-for="picture in bingPicturesPage.Content" :key="picture.Date">
        <a-card>
          <img :src="picture.Picture" slot="cover">
          <a-card-meta :title="picture.Title">
            <template slot="description">{{picture.Date}}</template>
          </a-card-meta>
        </a-card>
      </a-col>
    </a-row>
  </div>
</template>

<script>
  export default {
    name: "BingPicture",
    data() {
      return {
        bingPicturesPage: {}
      }
    },
    mounted() {
      this.fetchBingPictures();
    },
    methods: {
      fetchBingPictures() {
        this.api.picture.getBingPictures(1, 3).then(res => {
          this.bingPicturesPage = res.data;
        });
      }
    }
  }
</script>

<style scoped>

</style>