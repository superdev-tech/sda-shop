-- seed.sql : ข้อมูลตัวอย่างของร้าน (รันซ้ำกี่ครั้งก็ได้ เพราะใช้ ON CONFLICT DO NOTHING)

INSERT INTO categories (slug, name) VALUES
    ('coffee',      'กาแฟ'),
    ('tea',         'ชา'),
    ('snacks',      'ขนม'),
    ('accessories', 'อุปกรณ์')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO products (category_id, slug, name, description, price_satang, stock, image_url, is_featured)
SELECT c.id, p.slug, p.name, p.description, p.price_satang, p.stock, p.image_url, p.is_featured
FROM (VALUES
    ('coffee',      'doi-chang-arabica',  'กาแฟอาราบิก้าดอยช้าง 250 กรัม', 'คั่วกลาง หอมถั่วและช็อกโกแลต',           32000, 40, '/images/coffee.svg',    true),
    ('coffee',      'chiang-rai-espresso','เอสเพรสโซ่เบลนด์เชียงราย 500 กรัม', 'คั่วเข้ม เหมาะกับนมสด',                 45000, 25, '/images/coffee.svg',    false),
    ('coffee',      'cold-brew-bottle',   'โคลด์บรูว์ขวด 350 มล.',          'สกัดเย็น 18 ชั่วโมง',                     8900, 60, '/images/coffee.svg',    true),
    ('tea',         'thai-milk-tea-mix',  'ชาไทยปรุงสำเร็จ 400 กรัม',       'สูตรเข้มข้นแบบร้านดัง',                  15900, 80, '/images/tea.svg',       true),
    ('tea',         'jasmine-green-tea',  'ชาเขียวมะลิ 50 ซอง',             'ชาเขียวกลิ่นมะลิอ่อน ๆ',                 12900, 35, '/images/tea.svg',       false),
    ('tea',         'butterfly-pea-tea',  'ชาดอกอัญชัน 30 กรัม',            'ดอกอัญชันอบแห้ง ใส่มะนาวเปลี่ยนสี',       9900,  50, '/images/tea.svg',       false),
    ('snacks',      'durian-chips',       'ทุเรียนทอดกรอบ 100 กรัม',        'หมอนทองจันทบุรี ทอดกรอบ',                19900, 20, '/images/snack.svg',     true),
    ('snacks',      'mango-sticky-bar',   'ข้าวเหนียวมะม่วงอบแห้ง',          'ขนมทานเล่นรสมะม่วงน้ำดอกไม้',             5900,  100,'/images/snack.svg',     false),
    ('snacks',      'cashew-honey',       'เม็ดมะม่วงหิมพานต์อบน้ำผึ้ง 200 กรัม', 'เม็ดใหญ่ จากระนอง',                   22900, 30, '/images/snack.svg',     false),
    ('accessories', 'drip-kettle',        'กาดริปคอห่าน 600 มล.',           'สแตนเลส คุมสายน้ำได้แม่น',               69000, 10, '/images/accessory.svg', true),
    ('accessories', 'ceramic-mug',        'แก้วเซรามิกลายช้าง',             'ทำมือจากลำปาง ความจุ 300 มล.',           25000, 0,  '/images/accessory.svg', false),
    ('accessories', 'paper-filter',       'กระดาษกรองกาแฟ 100 แผ่น',        'ขนาด 02 ใช้กับดริปเปอร์ทรงกรวย',          7900,  120,'/images/accessory.svg', false)
) AS p(category_slug, slug, name, description, price_satang, stock, image_url, is_featured)
JOIN categories c ON c.slug = p.category_slug
ON CONFLICT (slug) DO NOTHING;

-- admin สำหรับทดสอบตอนพัฒนาเท่านั้น (รหัสผ่านจะบอกในคอร์สตอนทำส่วนหลังร้าน)
INSERT INTO users (email, password_hash, name, role) VALUES
    ('admin@sda-shop.dev', '$2a$10$5taLnOO98veWS5Ek.N/aQ.cLdv2PFlBU77AGxN2xXA5Ox63XRoPYy', 'แอดมินร้าน', 'admin')
ON CONFLICT (email) DO NOTHING;
