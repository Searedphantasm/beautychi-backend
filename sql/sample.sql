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
