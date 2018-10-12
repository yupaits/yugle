<template>
  <div>
    <breadcrumb  :items="['美图', 'Shot on OnePlus']"/>
    <div class="layout-content">
      <a-row :gutter="16">
        <a-col :span="8" v-for="picture in pictures" :key="picture.Date">
          <a-card class="mb-2">
            <img :src="picture.Picture" slot="cover">
            <a-card-meta :title="picture.Title">
              <template slot="description">{{picture.Date}} © {{picture.Author}}</template>
            </a-card-meta>
          </a-card>
        </a-col>
      </a-row>
      <div class="text-center">
        <a-pagination v-model="shotOnOnePlusPicturesPage.Page"
                      :total="shotOnOnePlusPicturesPage.Total"
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
    name: "ShotOnOnePlus",
    components: {Breadcrumb},
    data() {
      return {
        pictures: [],
        pageSize: 6,
        shotOnOnePlusPicturesPage: {
          page: 1,
          total: 0,
        },
      }
    },
    mounted() {
      this.fetchShotOnOnePlusPictures();
    },
    methods: {
      fetchShotOnOnePlusPictures() {
        this.$api.picture.getShotOnOnePlusPictures(this.shotOnOnePlusPicturesPage.page, this.pageSize).then(res => {
          this.pictures = res.data.Content;
          this.shotOnOnePlusPicturesPage.page = res.data.Page;
          this.shotOnOnePlusPicturesPage.total = res.data.Total;
        });
      },
      handlePagerChange(page, pageSize) {
        this.shotOnOnePlusPicturesPage.page = page;
        this.fetchShotOnOnePlusPictures();
      }
    }
  }
</script>

<style scoped>

</style>