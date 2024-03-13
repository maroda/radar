# TechScope

The beginning.

## Build

```zsh
go run . -list=TRad.csv > entries.json
cat index_hi.html entries.json index_low.html > index.html
cp index.html ../zalando/docs
```

## Issues

Not working yet because of this:


```
  entries: [							  entries: [
{							      |	      {
    "Quadrant": 0,					      |	        quadrant: 3,
    "Ring": 0,						      |	        ring: 0,
    "Label": "gRPC",					      |	        label: "AWS EMR",
    "Active": true,					      |	        active: false,
    "Link": ".",					      |	        link: "../data_processing/aws_emr.html",
    "Moved": 0						      |	        moved: 0
},							      |	      },
```

I think the labels are breaking it.
