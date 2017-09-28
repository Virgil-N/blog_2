var gulp = require("gulp");
var sass = require("gulp-sass-china");
var sourcemaps = require("gulp-sourcemaps");
var livereload = require('gulp-livereload');
var cssnano = require('gulp-cssnano');
var uglify = require('gulp-uglify');


gulp.task("css", function() {
    gulp.src("static/sass/**/*.scss")
            .pipe(sourcemaps.init())
            .pipe(sass())
            .pipe(cssnano())
            .pipe(sourcemaps.write("."))
            .pipe(gulp.dest("public/css"))
            .pipe(livereload());
});

gulp.task("js", function() {
  gulp.src("static/javascript/**/*.js")
        .pipe(sourcemaps.init())
        .pipe(uglify({mangle: false}))
        .pipe(sourcemaps.write("."))
        .pipe(gulp.dest("public/javascript"))
        .pipe(livereload())
});

gulp.task("vendor", function() {
  gulp.src("static/vendor/**/*")
        .pipe(gulp.dest("public/vendor"))
        .pipe(livereload())
});

gulp.task("img", function() {
  gulp.src("static/images/**/*")
        .pipe(gulp.dest("public/images"))
        .pipe(livereload())
});


gulp.task("fonts", function() {
  gulp.src("static/fonts/**/*")
        .pipe(gulp.dest("public/fonts"))
        .pipe(livereload())
});


gulp.task("watch", function() {
    livereload.listen();
    gulp.watch("static/sass/**/*.scss", ["css"]);
    gulp.watch("static/javascript/**/*.js", ["js"]);
    gulp.watch("static/vendor/**/*", ["vendor"]);
    gulp.watch("static/images/**/*", ["img"]);
    gulp.watch("static/fonts/**/*", ["fonts"]);
});