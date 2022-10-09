
const container = document.getElementById('mountNode');
if (container) {
	const width = container.clientWidth;
	const height = container.clientHeight;
	const graph = new G6.TreeGraph({
		animate: false,
		"container": container,
		"width": width,
		"height": height,
		"fitView": true,
		"fitViewPadding": [
			16,
			16,
			16,
			16
		],
		"layout": {
			'type': 'indented',
			'isHorizontal': true,
			// "type": "compactBox",
			'indent': 30,
			"direction": "LR",
			getHeight: function getHeight() {
				return 16;
			},
			getWidth: function getWidth() {
				return 16;
			},
			getVGap: function getVGap() {
				return 10;
			},
			getHGap: function getHGap() {
				return 100;
			},
		},
		'modes': {
			default: [
				{
					type: 'collapse-expand',
					animate: false,
					onChange: function onChange(item, collapsed) {
						const data = item.getModel();
						data.collapsed = collapsed;
						return true;
					},
				},
				'drag-canvas',
				'zoom-canvas',
			],
		},
		// "defaultNode": {
		// 	"type": "circle",
		// 	"linkPoints": {
		// 		"fill": "rgb(19, 30, 51)"
		// 	},
		// 	"labelCfg": {
		// 		"style": {
		// 			"style": {
		// 				"fill": "rgb(255, 255, 255)"
		// 			}
		// 		}
		// 	},
		// 	"size": 20
		// },
		// "defaultEdge": {
		// 	"type": "cubic-horizontal",
		// 	"style": {
		// 		"stroke": "rgb(79, 79, 79)"
		// 	}
		// },
	});


	graph.node(function (node) {
		return {
			size: 16,
			anchorPoints: [
				[0, 0.5],
				[1, 0.5],
			],
			label: node.lab,
			labelCfg: {
				offset: 5,
				position: node.children && node.children.length > 0 ? 'left' : 'right',
			},
			collapsed: node.collapsed,
			style: {
				fill: node.dir === true ? '#e6e56F' : '#3399FF',
				stroke: '#5B8FF9',
			},
		};
	});

	// graph.node((node) => {
	// 	const {
	// 		labelCfg = {}, icon = {}, linkPoints = {}, style = {}
	// 	} = node;
	// 	return {
	// 		...node,
	// 	}
	// })
	graph.edge((edge) => {
		return {
			type: 'cubic-horizontal',
			color: '#A3B1BF',
		}
	});
	// 读取数据
	graph.data(data);
	// 渲染图
	graph.render();
}	