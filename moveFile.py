import os
import re
import shutil


def migrate_to_subpackages():
    # 遍历当前目录下所有 .go 文件
    for filename in os.listdir('.'):
        if not filename.endswith('.go') or not os.path.isfile(filename):
            continue

        # 提取纯数字题号（例如：123.go → 123）
        match = re.match(r'^(\d+)\.go$', filename)
        if not match:
            print(f"跳过不符合命名规范的文件: {filename}")
            continue

        question_id = match.group(1)
        package_name = f"lc_{question_id}"
        target_dir = os.path.join(".", package_name)

        # 创建目标目录
        os.makedirs(target_dir, exist_ok=True)

        # 移动文件到目标目录
        src_path = filename
        dest_path = os.path.join(target_dir, filename)
        shutil.move(src_path, dest_path)
        print(f"移动文件: {src_path} → {dest_path}")

        # 修改文件内的包名
        with open(dest_path, 'r+', encoding='utf-8') as f:
            content = f.read()

            # 使用正则替换包声明（兼容多种写法）
            new_content = re.sub(
                r'^package\s+main\s*$',
                f'package {package_name}',
                content,
                flags=re.MULTILINE
            )

            # 回写文件
            f.seek(0)
            f.write(new_content)
            f.truncate()

        print(f"已更新包名为: {package_name}\n")


if __name__ == '__main__':
    migrate_to_subpackages()