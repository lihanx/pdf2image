# PDF2Image

## 简介
练手 + 帮女朋友解决pdf转图片的问题

## 注意
`Imagick`依赖`libmagickwand-dev`
但不能直接用于pdf处理，需要将`/etc/ImageMagick-6/policy.xml`中的
```
<policy domain="coder" rights="none" pattern="PDF" />
```
修改为
```
<policy domain="coder" rights="read|write" pattern="PDF" />
```
所以，本项目打包成docker镜像时会自动替换policy.xml，如果直接在机器的环境中运行，需要手动修改这一项。

## 其他
只是为了简单实现一个服务，代码没有经过雕琢和优化，对于分辨率和其他图片属性对设置可能需要重新调整代码。
如果有需要以后会再继续优化 *（这个flag我就立下了）* 。

以上。