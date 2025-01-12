INSERT INTO category (name, slug, description, image, image_key) VALUES
                                                                     ('الکترونیک', 'electronics', 'تمامی انواع دستگاه‌ها و گجت‌های الکترونیکی.', 'electronics.jpg', 'key_electronics'),
                                                                     ('مد و لباس', 'fashion', 'جدیدترین ترندها در پوشاک و لوازم جانبی.', 'fashion.jpg', 'key_fashion'),
                                                                     ('خانه و باغ', 'home-garden', 'مبلمان، دکور و لوازم باغبانی.', 'home_garden.jpg', 'key_home_garden'),
                                                                     ('ورزش', 'sports', 'لوازم و پوشاک ورزشی.', 'sports.jpg', 'key_sports');
INSERT INTO sub_category (parent_category_id, name, slug, description, image, image_key) VALUES
                                                                                             (1, 'تلفن‌های همراه', 'mobile-phones', 'جدیدترین تلفن‌های همراه از برندهای برتر.', 'mobile_phones.jpg', 'key_mobile_phones'),
                                                                                             (1, 'لپ‌تاپ‌ها', 'laptops', 'لپ‌تاپ‌های با عملکرد بالا برای کار و بازی.', 'laptops.jpg', 'key_laptops'),
                                                                                             (2, 'پوشاک مردانه', 'mens-clothing', 'پوشاک شیک برای مردان.', 'mens_clothing.jpg', 'key_mens_clothing'),
                                                                                             (2, 'پوشاک زنانه', 'womens-clothing', 'پوشاک مد روز برای زنان.', 'womens_clothing.jpg', 'key_womens_clothing');

INSERT INTO brand (name, slug, description, country, logo, logo_key, website_url) VALUES
                                                                                      ('اپل', 'apple', 'شرکت نوآور در زمینه فناوری.', 'آمریکا', 'apple_logo.png', 'key_apple', 'https://www.apple.com'),
                                                                                      ('نایک', 'nike', 'برند پیشرو در پوشاک و تجهیزات ورزشی.', 'آمریکا', 'nike_logo.png', 'key_nike', 'https://www.nike.com'),
                                                                                      ('سامسونگ', 'samsung', 'رهبر جهانی در الکترونیک مصرفی.', 'کره جنوبی', 'samsung_logo.png', 'key_samsung', 'https://www.samsung.com'),
                                                                                      ('آدیداس', 'adidas', 'مشهور به کفش‌ها و پوشاک ورزشی.', 'آلمان', 'adidas_logo.png', 'key_adidas', 'https://www.adidas.com');
INSERT INTO product (title, slug, description, poster, poster_key, price, category_id, brand_id, product_stock, product_discount_price, sub_category_id, consumer_guide, contact, status) VALUES
                                                                                                                                                                                              ('آیفون ۱۴', 'iphone-14', 'جدیدترین آیفون با ویژگی‌های پیشرفته.', 'iphone14.jpg', 'key_iphone14', 999, 1, 1, 50, 899, 1, 'راهنمای کاربر به صورت آنلاین موجود است.', '1234567890', 'Active'),
                                                                                                                                                                                              ('گلکسی S22', 'galaxy-s22', 'اسمارت‌فون با عملکرد بالا از سامسونگ.', 'galaxy_s22.jpg', 'key_galaxy_s22', 799, 1, 3, 30, 699, 1, 'راهنمای کاربر به صورت آنلاین موجود است.', '1234567891', 'Active'),
                                                                                                                                                                                              ('کفش نایک ایر مکس', 'nike-air-max', 'کفش‌های ورزشی راحت و شیک.', 'nike_air_max.jpg', 'key_nike_air_max', 150, 2, 2, 100, NULL, 3, 'به جدول سایز برای اندازه‌گیری مراجعه کنید.', '1234567892', 'Active'),
                                                                                                                                                                                              ('کفش آدیداس اولترا بوست', 'adidas-ultraboost', 'کفش‌های ورزشی با عملکرد بالا.', 'adidas_ultraboost.jpg', 'key_adidas_ultraboost', 180, 2, 4, 80, NULL, 3, 'به جدول سایز برای اندازه‌گیری مراجعه کنید.', '1234567893', 'Active');


INSERT INTO product_specifications (product_id, specs_title, specs_description) VALUES
                                                                                    (3, 'صفحه نمایش', 'صفحه نمایش 6.5 اینچی با کیفیت بالا'),
                                                                                    (2, 'پردازنده', 'پردازنده M1 با عملکرد فوق‌العاده'),
                                                                                    (3, 'تعداد صفحات', '300 صفحه');

INSERT INTO customer (username, first_name, last_name, email, phone) VALUES
                                                                         ('ali_karimi', 'علی', 'کریمی', 'ali.karimi@example.com', '09123456789'),
                                                                         ('sara_ahmadi', 'سارا', 'احمدی', 'sara.ahmadi@example.com', '09132345678'),
                                                                         ('reza_gholipour', 'رضا', 'قلی‌پور', 'reza.gholipour@example.com', '09143456789'),
                                                                         ('narges_mohammadi', 'نرگس', 'محمدی', 'narges.mohammadi@example.com', '09154567890'),
                                                                         ('kaveh_rahimi', 'کاوه', 'رحیمی', 'kaveh.rahimi@example.com', '09165678901');


INSERT INTO customer_address (customer_id, city, state, address, postal_code) VALUES
                                                                                  ('595b62d3-c94b-404b-b60a-cd86c68b6250', 'تهران', 'تهران', 'خیابان ولیعصر، کوچه یاس', '12345678901'),
                                                                                  ('68ac0587-8a76-45ba-bef7-3aa48e8c0768', 'اصفهان', 'اصفهان', 'خیابان چهارباغ بالا، نزدیک میدان نقش جهان', '23456789012'),
                                                                                  ('68ac0587-8a76-45ba-bef7-3aa48e8c0768', 'شیراز', 'فارس', 'بلوار زند، کوچه گلها', '34567890123'),
                                                                                  ('595b62d3-c94b-404b-b60a-cd86c68b6250', 'مشهد', 'خراسان رضوی', 'خیابان امام رضا، نزدیک حرم', '45678901234'),
                                                                                  ('8b6505d5-624e-4400-8160-7d37870a692a', 'تبریز', 'آذربایجان شرقی', 'خیابان ارتش جنوبی، کوچه شهید بهشتی', '56789012345');

