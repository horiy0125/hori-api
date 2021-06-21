class CreateQiitaPosts < ActiveRecord::Migration[6.1]
  def change
    create_table :qiita_posts do |t|
      t.string :title, null: false
      t.string :url, null: false
      t.string :thumbnail_url, null: false
      t.string :category, null: false

      t.timestamps
    end
  end
end
