# Handlebars Template

```
go build ;
./gtemplate g_template=your_template.tpl.handlebars \
            g_output=your_template.tpl.handlebars.out \
            g_type=handlebars \
            VAR_NAME="WALDINAR" \
            VAR_AGE=29 \
            VAR_ZIPCODE=04102000
```

# Go Text Template

```
go build ;
./gtemplate g_template=your_template.tpl.gotexttemplate \
            g_output=your_template.tpl.gotexttemplate.out \
            g_type=text_template \
            VAR_NAME="WALDINAR NETO" \
            VAR_AGE=23 \
            VAR_ZIPCODE=04102001
```
