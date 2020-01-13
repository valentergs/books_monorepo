var fs = require("fs");
var path = require("path");
var Walk = function(dir, done) {
  var results = [];
  var i = 0;
  var numberPattern = /\d+/g;

  fs.readdir(dir, function(err, list) {
    if (err) return done(err);

    (function next() {
      var file = list[i++];
      if (!file) return done(null, results);
      file = path.resolve(dir, file);
      fs.stat(file, function(err, stat) {
        if (stat && stat.isDirectory()) {
          Walk(file, function(err, res) {
            results = results.concat(res);
            next();
          });
        } else {
          file = file.match(numberPattern);
          if (file !== null) {
            results.push(file);
          }
          next();
        }
      });
    })();
  });
};

export default Walk;

Walk("../fotos", function(err, results) {
  if (err) throw err;
  console.log(results);
});
