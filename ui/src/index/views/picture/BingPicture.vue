<template>
  <div>
    <breadcrumb :items="['美图', '必应壁纸']"/>
    <div class="layout-content">
      <a-row :gutter="16">
        <a-col :span="8" v-for="picture in pictures" :key="picture.Date">
          <a-card class="mb-2">
            <img :src="picture.Picture" slot="cover">
            <a-card-meta :title="picture.Title">
              <template slot="description">{{picture.Date}}</template>
            </a-card-meta>
          </a-card>
        </a-col>
      </a-row>
      <div class="text-center">
        <a-pagination v-model="bingPicturesPage.Page"
                      :total="bingPicturesPage.Total"
                      :pageSize.sync="pageSize"
                      hideOnSinglePage
                      @change="handlePagerChange"></a-pagination>
      </div>
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
        pictures: [],
        pageSize: 6,
        bingPicturesPage: {
          page: 1,
          total: 0,
        },
      }
    },
    mounted() {
      this.fetchBingPictures();
    },
    methods: {
      fetchBingPictures() {
        this.$api.picture.getBingPictures(this.bingPicturesPage.page, this.pageSize).then(res => {
          this.pictures = res.data.Content;
          this.bingPicturesPage.page = res.data.Page;
          this.bingPicturesPage.total = res.data.Total;
        });
      },
      handlePagerChange(page, pageSize) {
        this.bingPicturesPage.page = page;
        this.fetchBingPictures();
      }
    }
  }
</script>

<style scoped>

</style>