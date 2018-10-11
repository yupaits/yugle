<template>
  <div>
    <breadcrumb :items="['美图', '必应壁纸']"/>
    <div class="layout-content">
      <a-row :gutter="16">
        <a-col :span="8" v-for="picture in bingPicturesPage.Content" :key="picture.Date">
          <a-card class="mb-2">
            <img :src="picture.Picture" slot="cover">
            <a-card-meta :title="picture.Title">
              <template slot="description">{{picture.Date}}</template>
            </a-card-meta>
          </a-card>
        </a-col>
      </a-row>
    </div>
  </div>
</template>

<script>
  import Breadcrumb from "../../../components/Breadcrumb";
  export default {
    name: "BingPicture",
    components: {Breadcrumb},
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
        this.$api.picture.getBingPictures(1, 6).then(res => {
          this.bingPicturesPage = res.data;
        });
      }
    }
  }
</script>

<style scoped>

</style>