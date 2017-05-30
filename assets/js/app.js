var MyVue = Vue.extend({
  delimiters: ['((', '))'],
  methods: {
    uploadFile: function(e){
      e.preventDefault()
      var files = e.target.files || e.dataTransfer.files
      var data = new FormData()
      var config = {
        headers: { 'content-type': 'multipart/form-data'}
      }
      data.append('file', files[0])
      axios.post('/api/files', data).then(function (response) {
        updateFileList()
      }).catch(function (error) {
        console.log(error)
      })
    },
    updateFileList: function(e) {
      axios.get('/api/files')
        .then(function (response) {
          filelist.files = response.data
        })
        .catch(function (error) {
          console.log(response)
        })
    }    
  }
})

var filelist = new MyVue({
  el: '#filelist',
  data: {
    files: []
  },
  methods: {
    deleteFile: function(e) {
      axios.delete('/api/files' + '/')
    }
  }
})
filelist.updateFileList()

var uploader = new MyVue({
  el: '#uploader'
})
