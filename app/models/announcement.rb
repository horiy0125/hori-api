class Announcement < ApplicationRecord

  validate :validate_link_url

  def validate_link_url
    errors.add(:link_url, 'は正しい形式で入力してください。') if link_url.match(%r{\Ahttp(s)?://}).nil? && link_url != ''
  end
end
