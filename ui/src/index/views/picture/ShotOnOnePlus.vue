<template>
  <div>
    <h2>Shot on OnePlus</h2>
    <a-row :gutter="16">
      <a-col :span="8" v-for="picture in shotOnOnePlusPicturesPage.Content" :key="picture.Date">
        <a-card class="mb-2">
          <img :src="picture.Picture" slot="cover">
          <a-card-meta :title="picture.Title">
            <template slot="description">{{picture.Date}} © {{picture.Author}}</template>
          </a-card-meta>
        </a-card>
      </a-col>
    </a-row>
  </div>
</template>

<script>
  export default {
    name: "ShotOnOnePlus",
    data() {
      return {
        shotOnOnePlusPicturesPage: {}
      }
    },
    mounted() {
      this.fetchShotOnOnePlusPictures();
    },
    methods: {
      fetchShotOnOnePlusPictures() {
        this.$api.picture.getShotOnOnePlusPictures(1, 6).then(res => {
          this.shotOnOnePlusPicturesPage = res.data;
        });
      }
    }
  }
</script>

<style scoped>

</style>