package schema

import (
	"github.com/GeertJohan/go.rice/embedded"
	"time"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "k.libsonnet",
		FileModTime: time.Unix(1523278962, 0),
		Content:     string("local k8s = import \"k8s.libsonnet\";\nlocal fn = {\n  mapContainers(f):: {\n    local podContainers = super.spec.template.spec.containers,\n    spec+: {\n      template+: {\n        spec+: {\n          containers: std.map(f, podContainers),\n        },\n      },\n    },\n  },\n  mapContainersWithName(names, f)::\n    local nameSet = if std.type(names) == \"array\" then std.set(names) else std.set([names]);\n    local inNameSet(name) = std.length(std.setInter(nameSet, std.set([name]))) > 0;\n\n    self.mapContainers(function(c) if std.objectHas(c, \"name\") && inNameSet(c.name) then f(c) else c),\n};\n\nk8s + {\n  apps:: k8s.apps + {\n    v1:: k8s.apps.v1 + {\n      daemonSet:: k8s.apps.v1.daemonSet + {\n        mapContainers(f):: fn.mapContainers(f),\n        mapContainersWithName(names, f):: fn.mapContainersWithName(names, f),\n      },\n      deployment:: k8s.apps.v1.deployment + {\n        mapContainers(f):: fn.mapContainers(f),\n        mapContainersWithName(names, f):: fn.mapContainersWithName(names, f),\n      },\n      replicaSet:: k8s.apps.v1.replicaSet + {\n        mapContainers(f):: fn.mapContainers(f),\n        mapContainersWithName(names, f):: fn.mapContainersWithName(names, f),\n      },\n      statefulSet:: k8s.apps.v1.statefulSet + {\n        mapContainers(f):: fn.mapContainers(f),\n        mapContainersWithName(names, f):: fn.mapContainersWithName(names, f),\n      },\n    },\n    v1beta1:: k8s.apps.v1beta1 + {\n      deployment:: k8s.apps.v1beta1.deployment + {\n        mapContainers(f):: fn.mapContainers(f),\n        mapContainersWithName(names, f):: fn.mapContainersWithName(names, f),\n      },\n      statefulSet:: k8s.apps.v1beta1.statefulSet + {\n        mapContainers(f):: fn.mapContainers(f),\n        mapContainersWithName(names, f):: fn.mapContainersWithName(names, f),\n      },\n    },\n    v1beta2:: k8s.apps.v1beta2 + {\n      daemonSet:: k8s.apps.v1beta2.daemonSet + {\n        mapContainers(f):: fn.mapContainers(f),\n        mapContainersWithName(names, f):: fn.mapContainersWithName(names, f),\n      },\n      deployment:: k8s.apps.v1beta2.deployment + {\n        mapContainers(f):: fn.mapContainers(f),\n        mapContainersWithName(names, f):: fn.mapContainersWithName(names, f),\n      },\n      replicaSet:: k8s.apps.v1beta2.replicaSet + {\n        mapContainers(f):: fn.mapContainers(f),\n        mapContainersWithName(names, f):: fn.mapContainersWithName(names, f),\n      },\n      statefulSet:: k8s.apps.v1beta2.statefulSet + {\n        mapContainers(f):: fn.mapContainers(f),\n        mapContainersWithName(names, f):: fn.mapContainersWithName(names, f),\n      },\n    },\n  },\n  batch:: k8s.batch + {\n    v1:: k8s.batch.v1 + {\n      job:: k8s.batch.v1.job + {\n        mapContainers(f):: fn.mapContainers(f),\n        mapContainersWithName(names, f):: fn.mapContainersWithName(names, f),\n      },\n    },\n    v1beta1:: k8s.batch.v1beta1 + {\n      cronJob:: k8s.batch.v1beta1.cronJob + {\n        mapContainers(f):: fn.mapContainers(f),\n        mapContainersWithName(names, f):: fn.mapContainersWithName(names, f),\n      },\n    },\n    v2alpha1:: k8s.batch.v2alpha1 + {\n      cronJob:: k8s.batch.v2alpha1.cronJob + {\n        mapContainers(f):: fn.mapContainers(f),\n        mapContainersWithName(names, f):: fn.mapContainersWithName(names, f),\n      },\n    },\n  },\n  core:: k8s.core + {\n    v1:: k8s.core.v1 + {\n      list:: {\n        new(items):: {\n          apiVersion: \"v1\",\n        } + {\n          kind: \"List\",\n        } + self.items(items),\n        items(items):: if std.type(items) == \"array\" then {items+: items} else {items+: [items]},\n      },\n      pod:: k8s.core.v1.pod + {\n        mapContainers(f):: fn.mapContainers(f),\n        mapContainersWithName(names, f):: fn.mapContainersWithName(names, f),\n      },\n      podTemplate:: k8s.core.v1.podTemplate + {\n        mapContainers(f):: fn.mapContainers(f),\n        mapContainersWithName(names, f):: fn.mapContainersWithName(names, f),\n      },\n      replicationController:: k8s.core.v1.replicationController + {\n        mapContainers(f):: fn.mapContainers(f),\n        mapContainersWithName(names, f):: fn.mapContainersWithName(names, f),\n      },\n    },\n  },\n  extensions:: k8s.extensions + {\n    v1beta1:: k8s.extensions.v1beta1 + {\n      daemonSet:: k8s.extensions.v1beta1.daemonSet + {\n        mapContainers(f):: fn.mapContainers(f),\n        mapContainersWithName(names, f):: fn.mapContainersWithName(names, f),\n      },\n      deployment:: k8s.extensions.v1beta1.deployment + {\n        mapContainers(f):: fn.mapContainers(f),\n        mapContainersWithName(names, f):: fn.mapContainersWithName(names, f),\n      },\n      replicaSet:: k8s.extensions.v1beta1.replicaSet + {\n        mapContainers(f):: fn.mapContainers(f),\n        mapContainersWithName(names, f):: fn.mapContainersWithName(names, f),\n      },\n    },\n  },\n}"),
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    "k8s.libsonnet",
		FileModTime: time.Unix(1523278962, 0),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1523279712, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "k.libsonnet"
			file3, // "k8s.libsonnet"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`assets`, &embedded.EmbeddedBox{
		Name: `assets`,
		Time: time.Unix(1523279712, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"k.libsonnet":   file2,
			"k8s.libsonnet": file3,
		},
	})
}