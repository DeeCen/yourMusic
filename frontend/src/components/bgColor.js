import { ColorfulBackgroundGenerator, ColorfulBackgroundLayer } from 'colorful-background-css-generator';

const rand = (r, e) => {
    return r + Math.round((e - r) * Math.random());
};

const buildLayer = (degree, s, l, posColor, posTransparency) => {
    return new ColorfulBackgroundLayer({
        degree: degree,
        h: rand(0, 360),
        s: s,
        l: l,
        posColor: posColor,
        posTransparency: posTransparency
    });
};

export const addBgColor = (id) => {
    let generator = new ColorfulBackgroundGenerator();
    const l1 = buildLayer(rand(0, 360), 0.9, 0.85, 100, 7);
    const l2 = buildLayer(rand(0, 300), 0.9, 0.7, 30, 90);
    const l3 = buildLayer(rand(0, 200), 0.9, 0.7, 10, 80);
    const l4 = buildLayer(rand(0, 100), 0.9, 0.65, 0, 70);
    const l5 = buildLayer(rand(0, 50), 0.9, 0.65, 0, 60);
    generator.addLayer(l1, undefined);
    generator.addLayer(l2, undefined);
    generator.addLayer(l3, undefined);
    generator.addLayer(l4, undefined);
    generator.addLayer(l5, undefined);
    generator.assignStyleToElementId(id);
};
