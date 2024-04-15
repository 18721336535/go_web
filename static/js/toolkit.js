var vue = new Vue({
    el: "#toolkit-app",
    data: {
        file: {id:"",fileName:"aaa",filePath:"",fileSize:"",lastUpdateTime:"",status:"",columns:[]},
        fileList: [],
        fileColumns: [],
        div_heads_message:""
    },
    methods: {
        findAll: function () {
            var _this = this;
            axios.get("/file/findAll").then(function (response) {
                _this.fileList = response.data;
                console.log(_this.fileList);
            }).catch(function (err) {
                console.log(err);
            });
        },
        findById: function (fileid) {
            var _this = this;
            axios.get("/file/findById", {
                params: {
                    id: fileid
                }
            }).then(function (response) {
              _this.file = response.data;
                $('#myModal').modal("show");
                $(document).ready(function () {
                var input_checkbox_content=document.getElementById("file-column-input-checkbox");
                var ii="x";
                var inputs="<div>check items:<br>";
                inputs+="<input type='checkbox' name='labelname' value='"+01+"' checked>"+empty+"</input>";
                inputs+="<input type='checkbox' name='labelname' value='"+02+"' checked>"+trailer+"</input>";
                inputs+="<input type='checkbox' name='labelname' value='"+03+"' checked>"+sum+"</input>";
                inputs+="<input type='checkbox' name='labelname' value='"+04+"' checked>"+duplicate_column+"</input>";
                inputs+="</div>";
                inputs+="<div>fileds to check:<br>:";
                for (var i = 1; i <= _this.file.columns.length; i++) {
                    inputs+="<input type='checkbox' name='labelname' value='"+i+"' checked>"+i+"."+_this.file.columns[i-1]+"</input>";
                    if ((i+1)%10 == 0) inputs+="<br>";
                }
                inputs+="</div>";
                input_checkbox_content.innerHTML=inputs;
                });
            }).catch(function (err) {
            });

        },
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
                  headsHtml = headsHtml+"</tr></table>";
                  _this.div_heads_message = headsHtml;
                  console.log(_this.fileList);
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
                  _this.fileList = response.data;
                  console.log(_this.fileList);
            }).catch(function (err) {
                  console.log(err);
            });
        },
        update: function (file) {
            var _this = this;
            axios.post("/file/update",_this.file).then(function (response) {
                _this.findAll();
            }).catch(function (err) {
            });
        }
    },

    created(){
        this.findAll();
    }
});