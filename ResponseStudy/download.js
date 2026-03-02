async downloadFile(row) {
    this.$http({
        method: 'post',
        url: 'file/upload',
        data:postData,
        responseType: "blob"
    }).then(res => {
        const _res = res.data
        let blob = new Blob([_res], {
            type: 'application/png'
        });
        let downloadElement = document.createElement("a");
        let href = window.URL.createObjectURL(blob); //创建下载的链接
        downloadElement.href = href;
        downloadElement.download = res.headers["fileName"]; //下载后文件名
        document.body.appendChild(downloadElement);
        downloadElement.click(); //点击下载
        document.body.removeChild(downloadElement); //下载完成移除元素
        window.URL.revokeObjectURL(href); //释放掉blob对象
    })}