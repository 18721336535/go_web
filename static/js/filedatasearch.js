var vue1 = new Vue({
    el: "#content-section",
    data: {
        file: {id:"",fileName:"aaa",filePath:"",fileSize:"",lastUpdateTime:"",status:"",columns:[]},
        fileColumns: [],
        fileDataRecords: [],
        div_heads_message:""
    },
    methods: {
        getFileHeads: function () {
            var _this = this;
            var filePath = $("#search-path-id").val();
            var field_key_value = $("#filed-search-key-value-id").val();
            let paramsDto = {filePath:filePath,fieldKeyValue:field_key_value}
            axios.post("/file/getFileHeads",paramsDto).then(function (response) {
                  var arr = [];
                  arr = response.data;
                  var headsHtml = "<table><tr>";
                  for(var i = 0; i < arr.length; i++){
                        headsHtml = headsHtml+"<td>"+arr[i]+"</td>";
                  }
                  headsHtml = headsHtml+"</tr><tr>";
                  for(var i = 1; i <= arr.length; i++){
                        headsHtml = headsHtml+"<td>"+i+"</td>";
                  }
                  headsHtml = headsHtml+"</tr></table>";
                  _this.div_heads_message = headsHtml;
            }).catch(function (err) {
                  console.log(err);
            });
        },

        findDataByFiledNameAndValue: function () {
            var _this = this;
            var filePath = $("#search-path-id").val();
            var field_key_value = $("#filed-search-key-value-id").val();
            let paramsDto = {filePath:filePath,fieldKeyValue:field_key_value}
            axios.post("/file/findDataByFiledNameAndValue",paramsDto).then(function (response) {
                  _this.fileDataRecords = response.data;
            }).catch(function (err) {
                  console.log(err);
            });
        }
    }
});